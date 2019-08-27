package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"mmrath.com/gobase/client/account"
	"mmrath.com/gobase/common/email"
	"mmrath.com/gobase/model"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}



func NewDB(cfg Config) (*model.DB, error) {
	return model.DBConn(cfg.DB)
}

func NewMailer(cfg Config) (email.Mailer, error) {
	return email.NewMailer(cfg.SMTP)
}

func NewNotifier(cfg Config, mailer email.Mailer) account.Notifier {
	return account.NewNotifier(cfg.Web.URL, mailer)
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(cfg Config, mux *chi.Mux) (*Server, error) {
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

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Print("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Print("shutting down server... reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Print("server gracefully stopped")
}
