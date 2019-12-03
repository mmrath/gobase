// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/mmrath/gobase/client/account"
	"github.com/mmrath/gobase/common/auth"
	"github.com/mmrath/gobase/common/email"
)

// Injectors from wire.go:

func BuildServer(config2 Config, mailer email.Mailer) (*Server, error) {
	notifier := NewNotifier(config2, mailer)
	db, err := NewDB(config2)
	if err != nil {
		return nil, err
	}
	service := account.NewService(notifier, db)
	handler := account.NewResource(service)
	jwtConfig := ProvideJWTConfig(config2)
	jwtService := auth.NewJWTService(jwtConfig)
	mux, err := NewMux(config2, handler, jwtService)
	if err != nil {
		return nil, err
	}
	server, err := NewServer(config2, mux)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

func ProvideJWTConfig(config2 Config) auth.JWTConfig {
	return config2.JWT
}

func ProvideSMTPConfig(config2 Config) email.SMTPConfig {
	return config2.SMTP
}