package app

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/mmrath/gobase/uaa-admin/pkg/account"
	"github.com/mmrath/gobase/common/config"
	"github.com/mmrath/gobase/model"
	"net/http"
	"os"
	"os/signal"
)

type App struct {
	httpServer *http.Server
}

func NewApp(profiles ...string) (*App, error) {
	cfg := Config{}
	err := config.LoadConfig(&cfg, "./resources", profiles...)
	if err != nil {
		return nil, err
	}

	db, err := model.DBConn(cfg.DB)

	if err != nil {
		return nil, err
	}

	roleService := account.NewRoleService(db)
	roleHandler := account.NewRoleHandler(roleService)

	httpHandler, err := HttpRouter(cfg.Web, roleHandler)

	if err != nil {
		return nil, err
	}

	httpServer := NewHttpServer(&cfg, httpHandler)
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
