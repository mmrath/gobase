package account

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/common/errors"
	"github.com/mmrath/gobase/model"
)

type Handler struct {
	service Service
}

func (h *Handler) Activate(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	err := h.service.Activate(key)

	if err != nil {
		errors.RenderError(w, r, err)
		return
	} else {
		render.Status(r, http.StatusOK)
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
		errors.RenderError(w, r, err)
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
		errors.RenderError(w, r, err)
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
		errors.RenderError(w, r, err)
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
		errors.RenderError(w, r, err)
		return
	} else {
		render.Status(r, http.StatusOK)
		render.JSON(w, r, user)
		return
	}

}
