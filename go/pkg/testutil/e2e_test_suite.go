package testutil

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"

	"github.com/mmrath/gobase/go/pkg/email"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	Handler     http.Handler
	server      *httptest.Server
	AppURL      string
	DB          *sql.DB
	EmailClient interface {
		GetLatestEmail(emailID string) *email.Message
	}
}

func (s *TestSuite) SetTestEnv() {
	// for randomness
	gofakeit.Seed(time.Now().UnixNano())
	portPrefix := os.Getenv("E2E_TEST_PORT_PREFIX")

	if portPrefix == "" {
		fmt.Println("E2E_TEST_PORT_PREFIX is not set. default to 40")
		portPrefix = "40"
	}

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", fmt.Sprintf("%s32", portPrefix))
	os.Setenv("DB_USERNAME", "app_user")
	os.Setenv("DB_PASSWORD", "password12")
	os.Setenv("DB_NAME", "appdb")
	os.Setenv("DB_SSLMODE", "disable")

	os.Setenv("SMTP_HOST", "localhost")
	os.Setenv("SMTP_PORT", fmt.Sprintf("%s13", portPrefix))
	os.Setenv("SMTP_FROM", "test@localhost")

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://app_user:password12@localhost:%s32/appdb?sslmode=disable", portPrefix)
		os.Setenv("DB_URL", dbURL)
	}
}

// SetupSuite setup at the beginning of test
func (s *TestSuite) SetupSuite() {
	portPrefix := os.Getenv("E2E_TEST_PORT_PREFIX")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		fmt.Printf("Failed to connect to DB %v", err)
		panic(err)
	}

	s.DB = db
	server := httptest.NewServer(s.Handler)
	appURL := server.URL
	mailURL := fmt.Sprintf("http://localhost:%s12", portPrefix)

	s.server = server
	s.AppURL = appURL
	s.EmailClient = NewEmailClient(mailURL)
}

// TearDownSuite teardown at the end of test
func (s *TestSuite) TearDownSuite() {
	s.server.Close()
	cleanDB(s.DB)
}

func (s *TestSuite) SetupTest() {
	// cleanDB(s.DB)
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

func MustExecStmt(db *sql.DB, stmt string, values ...interface{}) {
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
