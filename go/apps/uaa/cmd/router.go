package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mmrath/gobase/go/apps/uaa/internal/account"
	"github.com/mmrath/gobase/go/apps/uaa/internal/auth"
	"github.com/mmrath/gobase/go/apps/uaa/internal/config"
	"github.com/mmrath/gobase/go/apps/uaa/internal/static"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func BuildHttpServer(cfg *config.Config, sso *auth.Service, accountHandler *account.Handler) (*http.Server, error){
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(10 * time.Second))

	var contextPath string
	if cfg.Web.ContextPath != "" {
		contextPath = cfg.Web.ContextPath
	} else {
		contextPath = "/"
	}
	r.Route(contextPath, func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("pong"))
			if err != nil {
				log.Error().Err(err).Msg("error sending ping message")
			}
		})

		files, _ := static.WalkDirs("", true)
		log.Info().Interface("files", files).Msg("all")

		files, _ = static.WalkDirs("web", true)
		log.Info().Interface("web files", files).Msg("all")


		r.Get("/sso", auth.SsoGetHandler(static.HTTP))
		r.Post("/sso", auth.SsoPostHandler(sso))
		r.Post("/token", auth.TokenHandler(sso))
		r.Get("/logout", auth.LogoutHandler(sso))

		r.Group(func(r chi.Router) {

			r.Post("/account/sign-up", accountHandler.SignUp())
			r.Get("/account/activate", accountHandler.Activate())
			r.Post("/account/reset-password/init", accountHandler.InitPasswordReset())
			r.Post("/account/reset-password/finish", accountHandler.ResetPassword())
		})

		r.Get("/*", http.FileServer(static.HTTP).ServeHTTP)
	})

	r.NotFound(http.NotFound)

	srv := http.Server{
		Addr:    cfg.Web.Port,
		Handler: r,
	}
	return &srv,nil
}
