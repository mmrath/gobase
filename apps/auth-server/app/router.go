package app

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mmrath/gobase/apps/auth-server/internal/config"
	"github.com/mmrath/gobase/apps/auth-server/internal/oauth2"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func BuildHttpServer(cfg *config.Config) *http.Server {
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
		r.Route("/oauth2", func(r chi.Router) {
			r.HandleFunc("/auth", oauth2.AuthHandler())
			r.HandleFunc("/token", oauth2.TokenHandler())
			r.HandleFunc("/revoke", oauth2.RevokeHandler())
			r.HandleFunc("/introspect", oauth2.IntrospectionHandler())
		})
	})

	r.NotFound(http.NotFound)

	srv := http.Server{
		Addr:    cfg.Web.Port,
		Handler: r,
	}
	return &srv
}
