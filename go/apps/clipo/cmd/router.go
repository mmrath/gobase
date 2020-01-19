// Package api configures an http server for administration and application resources.
package cmd

import (
	"github.com/mmrath/gobase/go/apps/clipo/internal/config"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mmrath/gobase/go/apps/clipo/internal/account"
	"github.com/mmrath/gobase/go/pkg/auth"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// NewMux configures application resources and routes.
func NewMux(cfg config.Config,
	userHandler *account.Handler,
	jwtService auth.JWTService) (*chi.Mux, error) {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(10 * time.Second))
	//r.Use(log.NewStructuredLogger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if cfg.Web.CorsEnabled {
		r.Use(corsConfig().Handler)
	}

	r.Route("/api", func(r chi.Router) {
		// Protected routes
		r.Group(func(r chi.Router) {
			// Seek, verify and validate JWT tokens
			r.Use(jwtService.Verifier())

			// Handle valid / invalid tokens. In this example, we use
			// the provided authenticator middleware, but you can write your
			// own very easily, look at the Authenticator method in jwtauth.go
			// and tweak it, its not scary.
			r.Use(jwtService.Authenticator)

			r.Post("/account/change-password", userHandler.ChangePassword())

			r.Get("/account", userHandler.Account())
		})

		// Public routes
		r.Group(func(r chi.Router) {
			r.Post("/account/register", userHandler.Register())
			r.Get("/account/activate", userHandler.Activate())
			r.Post("/account/login", userHandler.Login(jwtService))
			r.Post("/account/logout", userHandler.Logout())
			r.Post("/account/reset-password/init", userHandler.InitPasswordReset())
			r.Post("/account/reset-password/finish", userHandler.ResetPassword())
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	client := "./public"
	r.Get("/*", spaHandler(client))

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

// spaHandler serves the public Single Page Application.
func spaHandler(publicDir string) http.HandlerFunc {
	handler := http.FileServer(http.Dir(publicDir))
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := path.Join(publicDir, "index.html")
		serviceWorker := path.Join(publicDir, "service-worker.js")

		requestedAsset := path.Join(publicDir, r.URL.Path)
		if strings.Contains(requestedAsset, "service-worker.js") {
			requestedAsset = serviceWorker
		}
		if _, err := os.Stat(requestedAsset); err != nil {
			http.ServeFile(w, r, indexPage)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
