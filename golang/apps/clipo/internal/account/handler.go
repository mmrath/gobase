package account

import (
	"fmt"
	"net/http"

	"github.com/mmrath/gobase/golang/pkg/errutil"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/golang/pkg/auth"
	"github.com/mmrath/gobase/golang/pkg/model"
)

var AuthTokenCookieName = "jwt"

type Handler struct {
	service *Service
}

func NewHandler(userService *Service) *Handler {
	return &Handler{service: userService}
}

func (h *Handler) Login(service auth.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.LoginRequest{}
		if err := render.DecodeJSON(r.Body, &data); err != nil {
			errutil.RenderError(w, r, err)
			return
		}
		user, err := h.service.Login(r.Context(), data)

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		var token string
		token, err = service.NewToken(&user)

		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:       AuthTokenCookieName,
			Value:      token,
			Path:       "/",
			RawExpires: "0",
			HttpOnly:   true,
		})

		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
		w.WriteHeader(http.StatusOK)
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
			render.JSON(w, r, errutil.Wrap(err, "failed to decode json"))
			return
		}

		user, err := h.service.Register(data)

		if err != nil {
			log.Error().Err(err).Msg("error during sign up")
			errutil.RenderError(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, &user)
	}
}

func (h *Handler) Activate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		err := h.service.Activate(key)
		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, struct{}{})
	}
}

func (h *Handler) GetProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := h.service.GetProfile(r.Context())
		if err != nil {
			errutil.RenderError(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, u)
	}
}

func (h *Handler) UpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.UserProfile{}

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		if err := h.service.UpdateProfile(r.Context(), data); err != nil {
			errutil.RenderError(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, nil)
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

		err := h.service.InitiatePasswordReset(data.Email)

		if err != nil {
			log.Error().Err(err).Msg("defaultError initiating password reset")
			errutil.RenderError(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
	}
}

func (h *Handler) ResetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := model.ResetPasswordRequest{}

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.service.ResetPassword(data)

		if err != nil {
			log.Error().Err(err).Msg("error initiating password reset")
			errutil.RenderError(w, r, err)
			return
		}

		render.Status(r, http.StatusOK)
	}
}

func (h *Handler) ChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := model.ChangePasswordRequest{}

		if err := render.DecodeJSON(r.Body, &data); err != nil {
			errutil.RenderError(w, r, err)
			return
		}

		err := h.service.ChangePassword(r.Context(), data)

		if err != nil {
			log.Error().Err(err).Msg("error changing password")
			errutil.RenderError(w, r, err)
			return
		}

		log.Info().Msg("password changed successfully")
		render.Status(r, http.StatusOK)
		render.PlainText(w, r, "Password changed successfully")
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
	}
}
