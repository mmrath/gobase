package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/model"
	"github.com/mmrath/gobase/uaa-server/pkg/config"
)

type App struct {
	httpServer *http.Server
}

func NewApp(profiles ...string) (*App, error) {
	cfg, err := config.LoadConfig("./resources", profiles...)
	if err != nil {
		return nil, err
	}

	_, err = model.DBConn(cfg.DB)

	if err != nil {
		return nil, err
	}

	httpHandler, err := HttpRouter(cfg)

	if err != nil {
		return nil, err
	}

	httpServer := NewHttpServer(cfg, httpHandler)
	return &App{httpServer: httpServer}, nil
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
