package account

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/error_util"
	"github.com/mmrath/gobase/common/template_util"
	"github.com/mmrath/gobase/model"
)

type Handler struct {
	service          Service
	templateRegistry *template_util.Registry
}

func (h *Handler) Activate(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	data := struct {
		Success bool
		Msg     string
	}{Success: false}

	err := h.service.Activate(key)
	if err != nil {
		data.Msg = err.Error()
		err := h.templateRegistry.Render(w, "activate.html", data)
		if err != nil {
			log.Error().Err(err).Msg("failed to write activation template")
			render.Status(r, http.StatusInternalServerError)
		}
		return
	} else {
		render.Status(r, http.StatusOK)
		data.Success = true
		err := h.templateRegistry.Render(w, "activation", data)

		if err != nil {
			log.Error().Err(err).Msg("failed to write activation template")
			render.Status(r, http.StatusInternalServerError)
		}

		render.PlainText(w, r, http.StatusText(http.StatusOK))
		return
	}
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) ResetPasswordFinish(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) PasswordResetInit(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

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
