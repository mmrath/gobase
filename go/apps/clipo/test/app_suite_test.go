package test

import (
	"context"
	"github.com/mmrath/gobase/go/apps/clipo/cmd"
	"github.com/mmrath/gobase/go/pkg/db"
	"log"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	db     *db.DB
	mailer *MockMailer
	server *httptest.Server
	cfg    cmd.Config
}

// SetupSuite setup at the beginning of test
func (s *TestSuite) SetupSuite() {

	_, file, _, _ := runtime.Caller(0)

	root := filepath.Dir(filepath.Dir(file))
	configRoot := filepath.Join(root, "resources")

	cfg := cmd.LoadConfig(configRoot, "test")

	mailer, err := NewMockMailer()
	require.NoError(s.T(), err)
	s.mailer = mailer
	db, err := db.Open(cfg.DB)
	require.NoError(s.T(), err)

	s.db = db
	s.cfg = cfg

	server, err := cmd.BuildServer(cfg, mailer)
	require.NoError(s.T(), err)
	s.server = httptest.NewServer(server.Handler)
}

// TearDownSuite teardown at the end of test
func (s *TestSuite) TearDownSuite() {
	defer s.server.Close()
}

func (s *TestSuite) SetupTest() {
	cleanDB(s.db)
	createTestUser(s.db)
}

func cleanDB(db *db.DB) {
	defer timeTrack(time.Now(), "truncate tables")
	stmts := []string{
		"TRUNCATE TABLE role_permission CASCADE",
		"TRUNCATE TABLE user_group_role CASCADE",
		"TRUNCATE TABLE user_group_user CASCADE",
		"TRUNCATE TABLE user_group CASCADE",
		"TRUNCATE TABLE user_role CASCADE",
		"TRUNCATE TABLE role CASCADE",
		"TRUNCATE TABLE user_credential CASCADE",
		"TRUNCATE TABLE permission CASCADE",
		"TRUNCATE TABLE user_account CASCADE",
		"TRUNCATE TABLE country CASCADE",
		"TRUNCATE TABLE currency CASCADE",
		"TRUNCATE TABLE timezone CASCADE",
		"TRUNCATE TABLE datetime_format CASCADE",
		"TRUNCATE TABLE date_format CASCADE",
		"TRUNCATE TABLE language CASCADE",
		"TRUNCATE TABLE notification CASCADE",
		"TRUNCATE TABLE notification_recipient CASCADE",
		"TRUNCATE TABLE notification_attachment CASCADE",
	}
	executeStmts(db, stmts)
}

func createTestUser(db *db.DB) {
	defer timeTrack(time.Now(), "create test user")
	stmts := []string{
		`INSERT INTO public.user_account(
	id, first_name, last_name, email, phone_number, active, created_at, created_by, updated_at, updated_by, version)
	VALUES (1, 'Test', 'Test', 'testuser@localhost', NULL, true, current_timestamp, 'test', current_timestamp, 'test', 1)`,
	}
	executeStmts(db, stmts)
}

func executeStmts(gormdb *db.DB, stmts []string) {
	_ = gormdb.RunInTx(context.Background(), func(tx *db.Tx) error {
		for _, stmt := range stmts {
			err := tx.Exec(stmt).Error
			if err != nil {
				log.Printf("Error executing statement %s", stmt)
				tx.Rollback()
				panic(err)
			}
		}
		return nil
	})
}

func execStmt(gormdb *db.DB, stmt string, values ...interface{}) error {
	err := gormdb.RunInTx(context.Background(), func(tx *db.Tx) error {
		err := tx.Exec(stmt, values).Error
		if err != nil {
			log.Printf("Error executing statement %s", stmt)
			return err
		}
		return nil
	})
	return err
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
