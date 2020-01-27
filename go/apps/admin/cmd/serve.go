package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func commandServe() *cobra.Command {
	return &cobra.Command{
		Use:     "serve [ env ]",
		Short:   "starts service with profiles",
		Long:    ``,
		Example: "serve dev",
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}
}

func serve(_ *cobra.Command, args []string) error {
	log.Info().Strs("arguments", args).Msg("starting application")

	app, err := NewApp()

	if err != nil {
		return err
	}

	app.Start()

	return nil
}
