package main

import (
	"fmt"
	"github.com/mmrath/gobase/go/apps/clipo/cmd"
	"github.com/mmrath/gobase/go/pkg/email"
	"github.com/mmrath/gobase/go/pkg/version"
	"os"
)

func main() {
	version.PrintVersion()

	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	cfg := cmd.LoadConfig()

	mailer, err := email.NewMailer(cfg.SMTP)
	if err != nil {
		fmt.Printf("Error creating mailer %v", err)
		os.Exit(1)
	}
	server, err := cmd.BuildServer(cfg, mailer)
	if err != nil {
		fmt.Printf("Exiting  %v", err)
		os.Exit(1)
	}
	server.Start()
}
