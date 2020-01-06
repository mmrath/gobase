package cmd

import (
	"context"
	"github.com/mmrath/gobase/go/apps/uaa/internal/auth"
	"github.com/mmrath/gobase/go/apps/uaa/internal/config"
	"github.com/mmrath/gobase/go/pkg/db"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
)

type App struct {
	httpServer *http.Server
}

func NewApp(configFiles ...string) (*App, error) {
	config := LoadConfig(configFiles...)
	sso, err := auth.NewSSO(config.Auth)
	if err != nil {
		return nil, err
	}

	db, err := db.Open(config.DB)
	if err != nil {
		return nil, err
	}
	srv, err := BuildHttpServer(config, sso)

	if err != nil {
		return nil, err
	}
	log.Info().Interface("config", config).Msg("App config")
	return &App{httpServer: srv}, nil
}

func LoadConfig(configFiles ...string) *config.Config {
	cfg, err := config.LoadConfig(configFiles...)
	if err != nil {
		log.Panic().Err(err).Msg("failed to load config")
		panic(err)
	}
	return cfg
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
