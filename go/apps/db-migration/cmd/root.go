package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"os"

	"github.com/mmrath/gobase/go/apps/db-migration/pkg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "db-migration",
	Short: "Upgrades DB",
	Long:  `Upgrades DB to the latest version by applying all the migrations incrementally`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.Upgrade()
		if err != nil {
			if err == migrate.ErrNoChange {
				fmt.Print("no changes detected")
				return
			}
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