package app

import (
	"context"
	"log"
	"mmrath.com/gobase/common/config"
	"net/http"
	"os"
	"os/signal"
)

type App struct {
	httpServer *http.Server
}

func New(profiles ...string) *App {
	cfg := &Config{}
	err := config.LoadConfig(&cfg, "./resources", profiles...)
	if err != nil {
		panic(err)
	}
	return &App{httpServer: NewServer(cfg)}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *App) Start() {
	log.Print("starting server...")
	go func() {
		if err := srv.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.httpServer.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Print("shutting down server... reason:", sig)

	if err := srv.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Print("server gracefully stopped")
}
