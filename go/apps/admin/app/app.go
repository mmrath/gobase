package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/go/apps/admin/internal/account"
	"github.com/mmrath/gobase/go/apps/admin/internal/config"
	"github.com/mmrath/gobase/go/pkg/db"
)

type App struct {
	httpServer *http.Server
}

func NewApp(profiles ...string) (*App, error) {
	cfg := config.Config{}
	err := config.LoadConfig(&cfg)
	if err != nil {
		return nil, err
	}

	db, err := db.Open(cfg.DB)

	if err != nil {
		return nil, err
	}

	roleService := account.NewRoleService(db)
	roleHandler := account.NewRoleHandler(roleService)

	httpHandler, err := NewHTTPRouter(cfg.Web, roleHandler)

	if err != nil {
		return nil, err
	}

	httpServer := NewHTTPServer(&cfg, httpHandler)
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Info().Interface("reason", sig).Msg("server shutting down")

	if err := srv.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Info().Msg("server stopped gracefully")
}
