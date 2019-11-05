// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/mmrath/gobase/common/email"
	"github.com/mmrath/gobase/uaa-server/internal/account"
	"net/http"
)

// Injectors from wire.go:

func BuildServer(mailer email.Mailer) *http.Server {
	config := LoadConfig()
	db := NewDB(config)
	notifier := account.NewNotifier(config, mailer)
	service := account.NewService(db, notifier)
	handler := account.NewHandler(service)
	appAppHandler := NewAppRouter(handler)
	httpHandler := HttpRouter(config, appAppHandler)
	server := NewHttpServer(config, httpHandler)
	return server
}
