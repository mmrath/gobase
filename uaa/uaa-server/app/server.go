package app

import (
	"net/http"
	"strings"

	"github.com/mmrath/gobase/uaa/uaa-server/internal/config"
)

func NewHttpServer(cfg *config.Config, handler http.Handler) *http.Server {
	var addr string
	port := cfg.Web.Port

	if strings.Contains(port, ":") {
		addr = port
	} else {
		addr = ":" + port
	}

	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return &srv
}
