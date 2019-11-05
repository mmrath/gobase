package account

import (
	"github.com/go-chi/chi"

	"github.com/mmrath/gobase/uaa-server/internal/config"
)

type Router struct {
	handler *Handler
}



func (h *Router) Register(r chi.Router, config *config.Config) {
	r.Route("/", func(r chi.Router) {
		r.Post("/signup", h.handler.SignUp)
		r.Post("/activate", h.handler.Activate)
		r.Post("/reset-password/init", h.handler.PasswordResetInit)
		r.Post("/reset-password/finish", h.handler.ResetPasswordFinish)
		r.Post("/change-password", h.handler.ChangePassword)
	})
}
