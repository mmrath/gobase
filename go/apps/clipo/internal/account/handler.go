package account

import (
	"fmt"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/model"
	"github.com/rs/zerolog/log"
)

var AuthTokenCookieName = "Token"

type Handler struct {
	Service *Service
}

func NewHandler(userService *Service) *Handler {
	return &Handler{Service: userService}
}

func (h *Handler) Login(service auth.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.LoginRequest{}
		if err := render.DecodeJSON(r.Body, &data); err != nil {
			render.JSON(w, r, err)
			return
		}
		user, err := h.Service.Login(data)

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		} else {
			var token string
			token, err = service.NewToken(&user)
			if err != nil {
				errutil.RenderError(w, r, err)
				return
			} else {
				render.Status(r, http.StatusOK)

				http.SetCookie(w, &http.Cookie{
					Name:       AuthTokenCookieName,
					Value:      token,
					Path:       "/",
					RawExpires: "0",
					HttpOnly:   true,
				})

				w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
				return
			}
		}
	}
}

func (h *Handler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:       AuthTokenCookieName,
			Value:      "",
			Path:       "/",
			RawExpires: "0",
			MaxAge:     -1, // Delete
			HttpOnly:   true,
		})
		render.Status(r, http.StatusOK)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.RegisterAccountRequest{}

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			render.JSON(w, r, err)
			return
		}

		user, err := h.Service.Register(data)

		if err != nil {
			log.Error().Err(err).Msg("error during sign up")
			errutil.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, &user)
			return
		}
	}
}

func (h *Handler) Activate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		err := h.Service.Activate(key)

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.PlainText(w, r, http.StatusText(http.StatusOK))
			return
		}
	}
}

func (h *Handler) GetProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusNotImplemented)
		return
	}
}

func (h *Handler) UpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusNotImplemented)
		return
	}
}

func (h *Handler) InitPasswordReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type InitPasswordResetRequest struct {
			Email string `json:"email"`
		}
		data := new(InitPasswordResetRequest)

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.Service.InitiatePasswordReset(data.Email)

		if err != nil {
			log.Error().Err(err).Msg("defaultError initiating password reset")
			errutil.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) ResetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.ResetPasswordRequest{}

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.Service.ResetPassword(data)

		if err != nil {
			log.Error().Err(err).Msg("error initiating password reset")
			errutil.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) ChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.ChangePasswordRequest{}

		if err := render.DecodeJSON(r.Body, data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.Service.ChangePassword(r.Context(), data)

		if err != nil {
			log.Error().Err(err).Msg("error changing password")
			errutil.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) Account() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			return
		}

		render.Status(r, http.StatusOK)
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.Raw))
		return
	}
}
