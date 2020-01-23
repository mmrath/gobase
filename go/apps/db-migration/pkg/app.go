package pkg

import (
	"github.com/golang-migrate/migrate/v4"

	// Loads PG driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	// Loads migration
	_ "github.com/mmrath/gobase/go/apps/db-migration/pkg/internal"
)

func Upgrade() error {
	m, err := buildMigration()
	if err != nil {
		return err
	}

	return doMigration(m.Up)
}

func Rollback() error {
	m, err := buildMigration()
	if err != nil {
		return err
	}

	return doMigration(m.Down)
}

func buildMigration() (*migrate.Migrate, error) {
	config := LoadConfig()
	migrationDir := "dir://" + config.MigrationDir

	return migrate.New(migrationDir, config.DB.URL())
}

func doMigration(action func() error) error {
	err := action()
	if err != nil {
		return err
	}

	return nil
}
