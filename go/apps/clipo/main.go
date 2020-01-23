//This needs installing tools locally to bin dir. Use `make tools`
//go:generate go-bindata -o internal/generated/assets.go -pkg generated -prefix "./resources/" "./resources/templates/..."

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

	app, err := cmd.BuildApp()
	if err != nil {
		fmt.Printf("Exiting  %v", err)
		os.Exit(1)
	}
	app.Start()
}
