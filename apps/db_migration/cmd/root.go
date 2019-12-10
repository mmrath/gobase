package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/mmrath/gobase/apps/db_migration/app"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fina",
	Short: "Upgrades DB",
	Long:  `Upgrades DB to the latest version by applying all the migrations incrementally`,
	Run: func(cmd *cobra.Command, args []string) {
		err := app.Upgrade()
		if err != nil {
			fmt.Printf("Error in upgrade step: %s", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
