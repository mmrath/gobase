package cmd

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mmrath/gobase/apps/clipo/internal/config"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func BuildHttpServer(cfg *config.Config) (*http.Server, error){
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("pong"))
			if err != nil {
				log.Error().Err(err).Msg("error sending ping message")
			}
		})
		r.Get("/*", spaHandler(cfg.Web.WebAppPath))

	})

	r.NotFound(http.NotFound)

	srv := http.Server{
		Addr:    cfg.Web.Port,
		Handler: r,
	}
	return &srv,nil
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
