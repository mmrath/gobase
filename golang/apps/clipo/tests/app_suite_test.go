package tests

import (
	"database/sql"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"

	"github.com/mmrath/gobase/golang/apps/clipo/cmd"
	"github.com/mmrath/gobase/golang/pkg/email"
	"github.com/mmrath/gobase/golang/pkg/testutil"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	app         *cmd.App
	server      *httptest.Server
	ClipoURL    string
	db          *sql.DB
	EmailClient interface {
		GetLatestEmail(emailID string) *email.Message
	}
}

// SetupSuite setup at the beginning of test
func (s *TestSuite) SetupSuite() {
	// for randomness
	gofakeit.Seed(time.Now().UnixNano())
	portPrefix := os.Getenv("E2E_TEST_PORT_PREFIX")

	if portPrefix == "" {
		fmt.Println("E2E_TEST_PORT_PREFIX is not set. default to 40")
		portPrefix = "40"
	}

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", fmt.Sprintf("%s32", portPrefix))
	os.Setenv("DB_USERNAME", "clipo")
	os.Setenv("DB_PASSWORD", "s3cr3t_1")
	os.Setenv("DB_NAME", "appdb")
	os.Setenv("DB_SSLMODE", "disable")

	os.Setenv("SMTP_HOST", "localhost")
	os.Setenv("SMTP_PORT", fmt.Sprintf("%s13", portPrefix))
	os.Setenv("SMTP_FROM", "test@localhost")

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://clipo:s3cr3t_1@localhost:%s32/appdb?sslmode=disable", portPrefix)
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		fmt.Printf("Failed to connect to DB %v", err)
		panic(err)
	}

	s.db = db

	app, err := cmd.BuildApp()

	if err != nil {
		panic(fmt.Errorf("failed to create server %w", err))
	}
	server := httptest.NewServer(app.Handler)

	clipoURL := server.URL
	mailURL := fmt.Sprintf("http://localhost:%s12", portPrefix)

	s.app = app
	s.server = server
	s.ClipoURL = clipoURL
	s.EmailClient = testutil.NewEmailClient(mailURL)
}

func apiPath(path string) string {
	return "/clipo/api" + path
}

// TearDownSuite teardown at the end of test
func (s *TestSuite) TearDownSuite() {
	s.server.Close()
	cleanDB(s.db)
}

func (s *TestSuite) SetupTest() {
	// cleanDB(s.db)
}

func cleanDB(db *sql.DB) {
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
		"TRUNCATE TABLE notification CASCADE",
		"TRUNCATE TABLE notification_recipient CASCADE",
		"TRUNCATE TABLE notification_attachment CASCADE",
	}
	executeStmts(db, stmts)
}

func executeStmts(db *sql.DB, stmts []string) {
	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("Error executing statement %s", stmt)
			panic(err)
		}
	}
}

func mustExecStmt(db *sql.DB, stmt string, values ...interface{}) {
	_, err := db.Exec(stmt, values...)
	if err != nil {
		log.Printf("Error executing statement %s, args %v", stmt, values)
		panic(err)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
