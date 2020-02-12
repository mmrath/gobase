package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/mmrath/gobase/apps/clipo/internal/config"
	"github.com/mmrath/gobase/apps/clipo/internal/templateutil"
	"github.com/mmrath/gobase/pkg/db"

	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/apps/clipo/internal/account"
	"github.com/mmrath/gobase/pkg/email"
)

// App provides an http.App.
type App struct {
	*http.Server
}

func NewDB(cfg config.Config) (*db.DB, error) {
	return db.Open(cfg.DB)
}

func NewNotifier(cfg config.Config, mailer email.Mailer, registry *templateutil.Registry) account.Notifier {
	return account.NewNotifier(cfg.AppDomainName, mailer, registry)
}

// NewApp creates and configures an APIServer serving all application routes.
func NewApp(cfg config.Config, mux http.Handler) (*App, error) {
	var addr string
	port := cfg.Web.Port

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &App{&srv}, nil
}

// Start runs ListenAndServe on the http.App with graceful shutdown.
func (srv *App) Start() {
	log.Print("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("listening on %s\n", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Print("shutting down server... reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Print("server gracefully stopped")
}
