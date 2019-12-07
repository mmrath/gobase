package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "admin",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
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
