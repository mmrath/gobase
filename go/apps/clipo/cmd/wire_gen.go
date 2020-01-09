package cmd

import (
	"github.com/mmrath/gobase/go/apps/clipo/internal/account"
	"github.com/mmrath/gobase/go/apps/clipo/internal/config"
	"github.com/mmrath/gobase/go/pkg/auth"
	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/errutil"
)

func BuildServer(config2 config.Config) (*Server, error) {

	mailer, err := email.NewMailer(config2.SMTP)
	if err != nil {
		return nil, errutil.Wrapf(err, "failed to build mailer")
	}
	notifier := NewNotifier(config2, mailer)
	db, err := NewDB(config2)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to create db connection")
	}
	service := account.NewService(notifier, db)
	handler := account.NewHandler(service)
	jwtService, err := auth.NewJWTService(config2.JWT)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to create JWT service")
	}
	mux, err := NewMux(config2, handler, jwtService)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to create http router")
	}
	server, err := NewServer(config2, mux)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to create server")
	}
	return server, nil
}
