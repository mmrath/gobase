package app

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/uaa-server/pkg/account"
	"github.com/mmrath/gobase/uaa-server/pkg/config"
	"github.com/mmrath/gobase/uaa-server/pkg/oauth2"
)

type appHandler struct {
	account *account.Handler
}

func NewAppRouter(account *account.Handler) *appHandler {
	return &appHandler{account: account}
}

func HttpRouter(cfg *config.Config, h *appHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(10 * time.Second))
	//r.Use(log.NewStructuredLogger(logger))
	//r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if cfg.Web.CorsEnabled {
		r.Use(corsConfig().Handler)
	}

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
			oauth2.RegisterHandlers(r,cfg)
		})
		r.Route("/account", func(r chi.Router) {
			r.Post("/signup", h.account.SignUp)
			r.Post("/activate", h.account.Activate)
			r.Post("/reset-password/init", h.account.PasswordResetInit)
			r.Post("/reset-password/finish", h.account.ResetPasswordFinish)
			r.Post("/change-password", h.account.ChangePassword)
		})
	})

	r.NotFound(http.NotFound)
	return r
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
