package test

import (
	"log"
	"github.com/mmrath/gobase/client/app"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/mmrath/gobase/model"
)

type TestSuite struct {
	suite.Suite
	db     *model.DB
	mailer *MockMailer
	server *httptest.Server
	cfg    app.Config
}

// SetupSuite setup at the beginning of test
func (s *TestSuite) SetupSuite() {

	_, file, _, _ := runtime.Caller(0)

	root := filepath.Dir(filepath.Dir(file))
	configRoot := filepath.Join(root, "resources")

	cfg := app.LoadConfig(configRoot, "test")

	mailer, err := NewMockMailer()
	require.NoError(s.T(), err)
	s.mailer = mailer
	db, err := model.DBConn(cfg.DB)
	require.NoError(s.T(), err)

	s.db = db
	s.cfg = cfg

	server, err := app.BuildServer(cfg, mailer)
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

func cleanDB(db *model.DB) {
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

func createTestUser(db *model.DB) {
	defer timeTrack(time.Now(), "create test user")
	stmts := []string{
		`INSERT INTO public.user_account(
	id, first_name, last_name, email, phone_number, active, created_at, created_by, updated_at, updated_by, version)
	VALUES (1, 'Test', 'Test', 'testuser@localhost', NULL, true, current_timestamp, 'test', current_timestamp, 'test', 1)`,
	}
	executeStmts(db, stmts)
}

func executeStmts(db *model.DB, stmts []string) {
	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error executing statement %s", stmt)
			panic(err)
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
