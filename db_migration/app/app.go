package app

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "mmrath.com/gobase/db_migration/internal"
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
	return migrate.New("dir://resources/migrations", config.DB.URL)
}

func doMigration(action func() error) error {
	err := action()
	if err != nil {
		return err
	}
	return nil
}
