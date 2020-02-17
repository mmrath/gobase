package cmd

import (
	"fmt"
	"os"

	"github.com/mmrath/gobase/golang/pkg/version"
)

func Main() {
	version.PrintVersion()
	app, err := BuildApp()
	if err != nil {
		fmt.Printf("Exiting  %v", err)
		os.Exit(1)
	}
	app.Start()
}
