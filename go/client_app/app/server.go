package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"mmrath.com/gobase/client/account"
	"mmrath.com/gobase/client/config"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"mmrath.com/gobase/pkg/email"
	"mmrath.com/gobase/pkg/model"
)

// Server provides an http.Server.
type Server struct {
	*http.Server
}

func LoadConfig() config.Config {
	return config.LoadConfig("./resources")
}

func NewDB(cfg config.Config) (*model.DB, error) {
	return model.DBConn(cfg.DB)
}

func NewMailer(cfg config.Config) (email.Mailer, error) {
	return email.NewMailer(cfg.SMTP)
}

func NewNotifier(cfg config.Config, mailer email.Mailer) account.Notifier {
	return account.NewNotifier(cfg.Server.URL, mailer)
}

// NewServer creates and configures an APIServer serving all application routes.
func NewServer(cfg config.Config, mux *chi.Mux) (*Server, error) {
	var addr string
	port := cfg.Server.Port

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
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
