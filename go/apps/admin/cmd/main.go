package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "admin",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Error().Err(err).Msg("failed to show help message")
			}
			os.Exit(2)
		},
	}
	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())

	return rootCmd
}

func Main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
