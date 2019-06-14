package main

import (
	"mmrath.com/gobase/client/app"
	"mmrath.com/gobase/pkg/email"
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
