package main

import (
	"fmt"
	"os"

	"github.com/mmrath/gobase/go/apps/clipo/cmd"
	"github.com/mmrath/gobase/go/pkg/version"
)

func main() {
	version.PrintVersion()

	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	cfg := cmd.LoadConfig()


	server, err := cmd.BuildServer(cfg)
	if err != nil {
		fmt.Printf("Exiting  %v", err)
		os.Exit(1)
	}
	server.Start()
}
