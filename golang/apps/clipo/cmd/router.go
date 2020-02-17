// Package api configures an http server for administration and application resources.
package cmd

import (
	"compress/flate"
	"github.com/mmrath/gobase/golang/pkg/health"
	"net/http"
	"time"

	"github.com/mmrath/gobase/golang/apps/clipo/internal/config"

	"github.com/mmrath/gobase/golang/apps/clipo/internal/account"
	"github.com/mmrath/gobase/golang/pkg/auth"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// NewMux configures application resources and routes.
func NewMux(cfg config.Config, userHandler *account.Handler, jwtService auth.JWTService) (*chi.Mux, error) {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.NewCompressor(flate.DefaultCompression).Handler())
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if cfg.Web.CorsEnabled {
		r.Use(corsConfig(cfg).Handler)
	}

	r.Route("/clipo/api", func(r chi.Router) {
		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(jwtService.Verifier())
			r.Use(jwtService.Authenticator)

			r.Post("/account/change-password", userHandler.ChangePassword())
			r.Get("/account", userHandler.Account())
		})

		// Public routes
		r.Group(func(r chi.Router) {
			r.Route("/account", func(r chi.Router) {
				r.Get("/activate", userHandler.Activate())
				r.Post("/register", userHandler.Register())
				r.Post("/login", userHandler.Login(jwtService))
				r.Post("/logout", userHandler.Logout())
				r.Post("/reset-password/init", userHandler.InitPasswordReset())
				r.Post("/reset-password/finish", userHandler.ResetPassword())
			})
			r.Get("/ping", health.PingHandlerFunc)
			r.HandleFunc("/*", http.NotFound)
		})


	})
	return r, nil
}

func corsConfig(cfg config.Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{cfg.AppDomainName},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}