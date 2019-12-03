package main

import (
	"github.com/mmrath/gobase/client/app"
	"github.com/mmrath/gobase/common/email"
)

func main() {

	cfg := app.LoadConfig()
	mailer, err := email.NewMailer(cfg.SMTP)
	if err != nil {
		panic(err)
	}
	server, err := app.BuildServer(cfg, mailer)
	if err != nil {
		panic(err)
	}
	server.Start()
}
