package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mmrath/gobase/apps/uaa/pkg/auth"
	"github.com/mmrath/gobase/apps/uaa/pkg/config"
	"github.com/mmrath/gobase/apps/uaa/static"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func BuildHttpServer(cfg *config.Config, sso *auth.SSOProvider) (*http.Server, error){
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
		r.Get("/sso", auth.SsoGetHandler(static.HTTP))
		r.Post("/sso", auth.SsoPostHandler(sso))
		r.Post("/auth_token", auth.AuthTokenHandler(sso))
		r.Get("/logout", auth.LogoutHandler(sso))
		r.Get("/*", http.FileServer(static.HTTP).ServeHTTP)
	})

	r.NotFound(http.NotFound)

	srv := http.Server{
		Addr:    cfg.Web.Port,
		Handler: r,
	}
	return &srv,nil
}
