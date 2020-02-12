package cmd

import (
	"github.com/mmrath/gobase/apps/clipo/internal/account"
	"github.com/mmrath/gobase/apps/clipo/internal/templateutil"
	"github.com/mmrath/gobase/pkg/auth"
	"github.com/mmrath/gobase/pkg/email"
	"github.com/mmrath/gobase/pkg/errutil"
)

func BuildApp() (*App, error) {
	config2 := LoadConfig()
	mailer, err := email.NewMailer(config2.SMTP)
	if err != nil {
		return nil, errutil.Wrapf(err, "failed to build mailer")
	}
	templateReg, err := templateutil.NewRegistry()

	if err != nil {
		return nil, errutil.Wrapf(err, "failed to build template registry")
	}

	notifier := NewNotifier(config2, mailer, templateReg)
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
	server, err := NewApp(config2, mux)
	if err != nil {
		return nil, errutil.Wrapf(err, "unable to create server")
	}
	return server, nil
}
