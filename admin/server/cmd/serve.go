package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func commandServe() *cobra.Command {
	return &cobra.Command{
		Use:     "serve [ config file ]",
		Short:   "Connect to the storage and begin serving requests.",
		Long:    ``,
		Example: "dex serve config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func serve(command *cobra.Command, strings []string) error {
	return nil
}
