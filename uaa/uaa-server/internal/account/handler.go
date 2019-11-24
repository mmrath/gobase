package account

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/error_util"
	"github.com/mmrath/gobase/common/template_util"
	"github.com/mmrath/gobase/model"
	"github.com/mmrath/gobase/uaa/uaa-server/internal/utils"
)

type Handler struct {
	service          Service
	templateRegistry *template_util.Registry
}

func (h *Handler) Activate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		err := h.service.Activate(key)

		if err != nil {
			h.renderError(w, r, "account/activation", err)
			return
		} else {
			h.renderSuccess(w, r, "account/activation", nil)
			return
		}
	}
}

func (h *Handler) ChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := new(model.ChangePasswordRequest)

		if err := render.DecodeJSON(r.Body, data); err != nil {
			render.JSON(w, r, err)
			return
		}
		err := h.service.ChangePassword(r.Context(), data)

		if err != nil {
			log.Error().Err(err).Msg("error changing password")
			error_util.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) ResetPasswordFinish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := new(model.ResetPasswordRequest)

		if err := render.DecodeJSON(r.Body, data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.service.ResetPassword(data)

		if err != nil {
			log.Error().Err(err).Msg("error initiating password reset")
			error_util.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) PasswordResetInit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type InitPasswordResetRequest struct {
			Email string `json:"email"`
		}
		data := new(InitPasswordResetRequest)

		if err := render.DecodeJSON(r.Body, data); err != nil {
			render.JSON(w, r, err)
			return
		}

		err := h.service.InitiatePasswordReset(data.Email)

		if err != nil {
			log.Error().Err(err).Msg("Error initiating password reset")
			error_util.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			return
		}
	}
}

func (h *Handler) SignUpForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := utils.NewTemplateRegistry()

		if err != nil {
			log.Error().Err(err).Msg("filed to load templates")
		}

		err = t.RenderTemplate(w, "account/sign-up-form.html", "")
		if err != nil {
			log.Error().Err(err).Msg("filed to render template")
		}
	}
}

func (h *Handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &model.SignUpRequest{}

		if err := render.DecodeJSON(r.Body, data); err != nil {
			render.JSON(w, r, err)
			return
		}

		user, err := h.service.SignUp(data)

		if err != nil {
			log.Error().Err(err).Msg("error during sign up")
			error_util.RenderError(w, r, err)
			return
		} else {
			render.Status(r, http.StatusOK)
			render.JSON(w, r, user)
			return
		}
	}
}

func (h *Handler) renderSuccess(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	err := h.templateRegistry.RenderHttp(w, templateName, data)
	h.handleInternalError(nil, nil, err)
}

func (h *Handler) renderError(w http.ResponseWriter, r *http.Request, templateName string, err error) {
	data := map[string]interface{}{
		"success": false,
	}
	err = h.templateRegistry.RenderHttp(w, templateName, data)
	h.handleInternalError(r, w, err)
}

func (h *Handler) handleInternalError(r *http.Request, w http.ResponseWriter, e error) {
	render.Status(r, http.StatusInternalServerError)
	log.Error().Err(e).Msg("internal error")

}
