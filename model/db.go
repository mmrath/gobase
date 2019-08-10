package model

import (
	"context"
	"log"

	"github.com/go-pg/pg"
)

type DBConfig struct {
	URL   string `yaml:"url"`
	Debug bool   `yaml:"debug"`
}

// DBConn returns a postgres connection pool.
func DBConn(config DBConfig) (*DB, error) {

	opts, err := pg.ParseURL(config.URL)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opts)
	if err := checkConn(db); err != nil {
		return nil, err
	}

	if config.Debug {
		db.AddQueryHook(&logSQL{})
	}

	return &DB{db}, nil
}

type logSQL struct{}

func (l *logSQL) BeforeQuery(e *pg.QueryEvent) {}

func (l *logSQL) AfterQuery(e *pg.QueryEvent) {
	query, err := e.FormattedQuery()
	if err != nil {
		panic(err)
	}
	log.Println(query)
}

func checkConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

type DB struct {
	*pg.DB
}

type Tx struct {
	*pg.Tx
	ctx context.Context
}

func IsNoDataFound(err error) bool {
	return err == pg.ErrNoRows
}

func (tx *Tx) cleanUp() {
	if err := recover(); err != nil {
		_ = tx.Rollback()
		panic(err)
	}
}

func (db *DB) RunTx(fn func(tx *Tx) error) error {
	pgTx, err := db.Begin()
	if err != nil {
		return err
	}
	tx := &Tx{pgTx, context.Background()}

	defer tx.cleanUp()

	if err := fn(tx); err != nil {
		_ = tx.Commit()
		return err
	}
	return tx.Commit()
}

func (db *DB) Tx(ctx context.Context, fn func(tx *Tx) error) error {
	pgTx, err := db.Begin()
	if err != nil {
		return err
	}
	tx := &Tx{pgTx, ctx}

	defer tx.cleanUp()

	if err := fn(tx); err != nil {
		_ = tx.Commit()
		return err
	}
	return tx.Commit()
}

func (tx *Tx) UserDao() UserDao {
	return newUserDao(tx)
}

func (tx *Tx) UserCredentialDao() UserCredentialDao {
	return newUserCredentialDao(tx)
}
