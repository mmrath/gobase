package db

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
)

type txKeyType string

var txKey txKeyType

type Config struct {
	Host     string
	Port     int `default:"5432"`
	Username string
	Password string
	Name     string
	SSLMode  string `default:"enabled"`
	Debug    bool   `default:"false" yaml:"debug"`
}

// DBConn returns a postgres connection pool.
func Open(cfg Config) (*DB, error) {
	log.Info().Msg("trying to connect to db")
	db, err := gorm.Open("postgres", cfg.URL())

	if err != nil {
		return nil, eris.Wrapf(err, "failed to open connection")
	}
	log.Info().Msg("successfully connected to db")
	if cfg.Debug {
		db = db.Debug()
	}
	return &DB{db}, nil
}

func (c Config) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.Name, c.SSLMode)
}

type DB struct {
	gorm *gorm.DB
}

func IsNoDataFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func (db *DB) RunInTx(ctx context.Context, fn func(tx *Tx) error) error {
	gormTx := db.gorm.BeginTx(ctx, nil)
	if gormTx.Error != nil {
		return eris.Wrapf(gormTx.Error, "failed to begin db transaction")
	}
	tx := &Tx{gormTx}

	defer tx.cleanUp()

	if err := fn(tx); err != nil {
		_ = tx.Commit()
		return err
	}
	return tx.Commit().Error
}

type Tx struct {
	*gorm.DB
}

func (tx *Tx) cleanUp() {
	if err := recover(); err != nil {
		_ = tx.Rollback()
		panic(err)
	}
}