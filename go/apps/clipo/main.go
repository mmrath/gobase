package main

import (
	"github.com/mmrath/gobase/go/apps/clipo/cmd"
	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/version"
)

func main() {
	version.PrintVersion()

	cfg := cmd.LoadConfig()
	mailer, err := email.NewMailer(cfg.SMTP)
	if err != nil {
		panic(err)
	}
	server, err := cmd.BuildServer(cfg, mailer)
	if err != nil {
		panic(err)
	}
	server.Start()
}
