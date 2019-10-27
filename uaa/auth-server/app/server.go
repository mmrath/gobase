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

	"github.com/mmrath/gobase/common/log"
	"github.com/mmrath/gobase/uaa-server/pkg/config"
	"github.com/mmrath/gobase/uaa-server/pkg/handlers"
)

func NewHttpServer(cfg *config.Config, handler http.Handler) *http.Server {
	var addr string
	port := cfg.Web.Port

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &srv
}

func HttpRouter(cfg *config.Config) (http.Handler, error) {
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
	if cfg.Web.ContextPath != ""{
		contextPath = cfg.Web.ContextPath
	}else {
		contextPath = "/"
	}
	r.Route(contextPath, func(r chi.Router) {

		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("pong"))
			if err != nil{
				log.Errorf("error sending ping message %v", err)
			}
		})

		handlers.RegisterRoutes(r, cfg)

	})

	r.NotFound(http.NotFound)
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
