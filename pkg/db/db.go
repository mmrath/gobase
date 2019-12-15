package db

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mmrath/gobase/pkg/error_util"
)

type txKeyType string

var txKey txKeyType

type Config struct {
	URL   string `yaml:"url"`
	Debug bool   `yaml:"debug"`
}

// DBConn returns a postgres connection pool.
func Open(cfg Config) (*DB, error) {
	db, err := gorm.Open("postgres", cfg.URL)
	if err != nil {
		return nil, error_util.NewInternal(err, "failed to open connection")
	}
	if cfg.Debug {
		db = db.Debug()
	}
	return &DB{db}, nil
}

type DB struct {
	gorm *gorm.DB
}

func IsNoDataFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func (db *DB) Tx(ctx context.Context, fn func(tx *Tx) error) error {
	gormTx := db.gorm.BeginTx(ctx, nil)
	if gormTx.Error != nil {
		return error_util.NewInternal(gormTx.Error, "failed to begin db transaction")
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
