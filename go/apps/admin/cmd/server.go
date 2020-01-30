package cmd

import (
	"net/http"
	"strings"

	"github.com/mmrath/gobase/go/apps/admin/internal/config"
)

func NewHTTPServer(cfg *config.Config, handler http.Handler) *http.Server {
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
