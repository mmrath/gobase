package cmd

import (
	"compress/flate"
	"github.com/mmrath/gobase/golang/pkg/health"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mmrath/gobase/golang/apps/oppo/internal/account"
	"github.com/mmrath/gobase/golang/apps/oppo/internal/config"
)

func NewHTTPRouter(webConfig config.WebConfig, rh *account.RoleHandler, uh *account.UserHandler) (http.Handler, error) {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.NewCompressor(flate.DefaultCompression).Handler())
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if webConfig.CorsEnabled {
		r.Use(corsConfig().Handler)
	}

	r.Route("/oppo/api", func(r chi.Router) {
		// Protected routes
		r.Group(func(r chi.Router) {
			r.Route("/role", func(r chi.Router) {
				r.Get("/:id", rh.FindRole)
				r.Post("/", rh.CreateRole)
				r.Put("/:id", rh.UpdateRole)
			})

			r.Route("/account", func(r chi.Router) {
				r.Get("/:id", uh.FindUser)
				r.Post("/", uh.CreateUser)
				r.Put("/:id", uh.UpdateUser)
			})
			r.Get("/ping", health.PingHandlerFunc)
			r.HandleFunc("/*", http.NotFound)
		})
	})

	return r, nil
}

func corsConfig() *cors.Cors {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}
