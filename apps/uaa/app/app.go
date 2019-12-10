package app

import (
	"context"
	 "github.com/mmrath/gobase/pkg/db"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/apps/uaa/internal/config"
)

type App struct {
	httpServer *http.Server
}

func NewApp(configFiles ...string) (*App, error) {
	config := LoadConfig(configFiles...)
	log.Info().Interface("config", config).Msg("App config")
	httpServer := BuildServer(config, nil)
	return &App{httpServer: httpServer}, nil
}

func LoadConfig(configFiles ...string) *config.Config {

	cfg, err := config.LoadConfig(configFiles...)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
		panic(err)
	}
	return cfg
}

func NewDB(cfg *config.Config) *db.DB {
	db, err := db.Open(cfg.DB)
	if err != nil {
		panic(err)
	}
	return db
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *App) Start() {
	log.Info().Msg("server starting")
	go func() {
		if err := srv.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Info().Interface("address", srv.httpServer.Addr).Msg("server started")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Info().Interface("reason", sig).Msg("server shutting down")

	if err := srv.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Info().Msg("server stopped gracefully")
}
