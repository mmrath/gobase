package tests

import (
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/mmrath/gobase/go/pkg/email"
	"log"
	"os"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	ClipoURL    string
	db          *sql.DB
	EmailClient interface {
		GetLatestEmail(emailId string) *email.Message
	}
}

// SetupSuite setup at the beginning of test
func (s *TestSuite) SetupSuite() {
	// for randomness
	gofakeit.Seed(time.Now().UnixNano())

	portPrefix := os.Getenv("ENV_PORT_PREFIX")

	if portPrefix == "" {
		fmt.Println("ENV_PORT_PREFIX is not set")
		panic("ENV_PORT_PREFIX is not set")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://app_user:password12@localhost:%s32/appdb?sslmode=disable", portPrefix)
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		fmt.Printf("Failed to connect to DB %v", err)
		panic(err)
	}

	s.db = db

	clipoURL := fmt.Sprintf("http://localhost:%s10", portPrefix)
	mailURL := fmt.Sprintf("http://localhost:%s12", portPrefix)

	s.ClipoURL = clipoURL
	s.EmailClient = NewTestEmailClient(mailURL)
}

// TearDownSuite teardown at the end of test
func (s *TestSuite) TearDownSuite() {
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

func createTestUser(db *sql.DB) {
	defer timeTrack(time.Now(), "create test user")
	stmts := []string{
		`INSERT INTO public.user_account(
	id, first_name, last_name, email, phone_number, active, created_at, created_by, updated_at, updated_by, version)
	VALUES (1, 'Test', 'Test', 'testuser@localhost', NULL, true, current_timestamp, 'test', current_timestamp, 'test', 1)`,
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

func (s *TestSuite) PopLastMessage(email string) MailMessage {
	return MailMessage{}
}
