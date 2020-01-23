// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/mmrath/gobase/go/apps/db-migration/pkg"

	// Loads driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	"github.com/spf13/cobra"
)

// rollback_migration represents the upgradeDb command
var rollbackMigration = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback DB to the previous version",
	Long:  `Rollback DB to the previous version by applying down migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.Rollback()
		if err != nil {
			fmt.Printf("Error in rollback step: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rollbackMigration)
}
