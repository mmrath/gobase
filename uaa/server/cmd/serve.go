package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mmrath.com/gobase/uaa/app"
	"os"
)

func commandServe() *cobra.Command {
	return &cobra.Command{
		Use:     "serve [ config file ]",
		Short:   "Connect to the storage and begin serving requests.",
		Long:    ``,
		Example: "serve config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func serve(command *cobra.Command, strings []string) error {
	app.New().Start()
	return nil
}
