package tests

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/mmrath/gobase/golang/apps/clipo/cmd"
	"github.com/mmrath/gobase/golang/pkg/crypto"
	"github.com/mmrath/gobase/golang/pkg/testutil"

	"github.com/mmrath/gobase/golang/pkg/model"
)

type AccountTestSuite struct {
	testutil.TestSuite
}

func (s *AccountTestSuite) SetupSuite() {
	s.TestSuite.SetTestEnv()
	app, err := cmd.BuildApp()
	if err != nil {
		panic(err)
	}
	s.Handler = app.Handler
	s.TestSuite.SetupSuite()
}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (s *AccountTestSuite) TestPing() {
	resp, err := http.Get(s.AppURL + "/clipo/api/ping")
	require.NoError(s.T(), err)
	require.Equal(s.T(), 200, resp.StatusCode)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	require.NoError(s.T(), err)
	contents := buf.String()
	require.Equal(s.T(), "pong", contents)
}

func (s *AccountTestSuite) TestRegisterActivateAndLogin() {
	he := httpexpect.New(s.T(), s.AppURL)
	testEmail := gofakeit.Email()
	testPassword := gofakeit.Password(true, true, true, true, true, 8)
	registerRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     testEmail,
		"password":  testPassword,
	}
	fmt.Printf("registering with data %v\n", registerRequest)
	he.POST(apiPath("/account/register")).
		WithJSON(registerRequest).
		Expect().
		Status(http.StatusOK)

	msg := s.EmailClient.GetLatestEmail(testEmail)
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Activate your account", msg.Subject)
	require.Equal(s.T(), registerRequest["email"], msg.To[0].Address)

	resp := he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: testPassword}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	resp.JSON().Path("$.errors[0]").Equal("user is not activated")

	re := regexp.MustCompile(`/account/activate\?key=([0-9a-f\-]+)`)
	key := re.FindStringSubmatch(msg.HTML)[1]
	he.GET(apiPath("/account/activate")).
		WithQuery("key", key).
		Expect().
		Status(http.StatusOK)

	resp = he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: testPassword}).
		Expect()
	resp.Status(http.StatusOK)
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)
}

func (s *AccountTestSuite) TestRegisterWithInvalidEmail() {
	he := httpexpect.New(s.T(), s.AppURL)
	registerRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     "invalid_email",
		"password":  gofakeit.Password(true, true, true, true, true, 8),
	}
	resp := he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.fieldErrors[0].field").Equal("email")
	resp.JSON().Path("$.fieldErrors[0].message").Equal("email must be a valid email address")
}

func (s *AccountTestSuite) TestRegisterWithDuplicateEmail() {
	he := httpexpect.New(s.T(), s.AppURL)
	registerRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true, true, true, true, true, 8),
	}
	resp := he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusOK)

	// 2nd Request
	resp = he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusBadRequest)

	resp.JSON().Path("$.fieldErrors[0].field").Equal("email")
	resp.JSON().Path("$.fieldErrors[0].message").Equal("email already registered")
}

func (s *AccountTestSuite) TestRegisterWithInvalidPassword() {
	he := httpexpect.New(s.T(), s.AppURL)
	registerRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true, true, true, true, false, 5), /// too short
	}
	// 2nd Request
	resp := he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusBadRequest)

	resp.JSON().Path("$.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.fieldErrors[0].message").Equal("password must be at least 6 characters in length")
}

func (s *AccountTestSuite) TestRegisterWithLongPassword() {
	he := httpexpect.New(s.T(), s.AppURL)
	registerRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true, true, true, true, true, 20),
	}
	// 2nd Request
	resp := he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusOK)
}

func (s *AccountTestSuite) TestRegisterWithTooLongPassword() {
	he := httpexpect.New(s.T(), s.AppURL)
	registerRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true, true, true, true, true, 32),
	}
	// 2nd Request
	resp := he.POST(apiPath("/account/register")).WithJSON(registerRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.fieldErrors[0].message").Equal("password must be a maximum of 20 characters in length")
}

func (s *AccountTestSuite) TestActivateWithWrongKey() {
	he := httpexpect.New(s.T(), s.AppURL)
	resp := he.GET(apiPath("/account/activate")).
		WithQuery("key", "wrong-key").
		Expect().
		Status(http.StatusBadRequest)
	resp.JSON().Path("$.errors[0]").Equal("invalid activation token")
}

func (s *AccountTestSuite) TestResetPassword() {
	testEmail := gofakeit.Email()
	he := httpexpect.New(s.T(), s.AppURL)
	initResetRequest := map[string]interface{}{
		"email": testEmail,
	}
	s.createUser(testEmail, gofakeit.Person().LastName)
	resp := he.POST(apiPath("/account/reset-password/init")).
		WithJSON(initResetRequest).Expect()
	resp.Status(http.StatusOK)

	// Login should throw an error now
	msg := s.EmailClient.GetLatestEmail(testEmail)
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Reset password", msg.Subject)
	require.Contains(s.T(), msg.To[0].Address, initResetRequest["email"])

	re := regexp.MustCompile(`/account/reset-password\?key=([0-9a-f\-]+)`)
	key := re.FindStringSubmatch(msg.HTML)[1]

	resetRequest := map[string]interface{}{
		"resetToken":  key,
		"newPassword": "Secret123",
	}

	he.POST(apiPath("/account/reset-password/finish")).
		WithJSON(resetRequest).
		Expect().
		Status(http.StatusOK)

	resp = he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: "Secret123"}).
		Expect()

	resp.Status(http.StatusOK)
}

func (s *AccountTestSuite) TestWithWrongUsername() {
	testEmail := gofakeit.Email()
	he := httpexpect.New(s.T(), s.AppURL)

	resp := he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: "Secret123"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.JSON().Path("$.errors[0]").Equal("invalid email or password")
}

func (s *AccountTestSuite) TestWithWrongPassword() {
	testEmail := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, true, 8)
	s.createUser(testEmail, password)
	defer s.deleteUser(testEmail)

	he := httpexpect.New(s.T(), s.AppURL)

	resp := he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: "incorrectPassword"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.JSON().Path("$.errors[0]").Equal("invalid email or password")
}

func (s *AccountTestSuite) TestChangePassword() {
	testEmail := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, true, 8)
	newPassword := gofakeit.Password(true, true, true, true, true, 10)
	s.createUser(testEmail, password)
	defer s.deleteUser(testEmail)

	he := httpexpect.New(s.T(), s.AppURL)

	resp := he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusOK)

	fmt.Printf("Cookies %v", resp.Cookies().Iter())
	jwtCookie := resp.Cookie("jwt").Value().Raw()
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)

	resp = he.POST(apiPath("/account/change-password")).
		WithJSON(model.ChangePasswordRequest{CurrentPassword: password, NewPassword: newPassword}).
		WithHeader("Authorization", "Bearer "+token).
		WithCookie("jwt", jwtCookie).
		Expect()
	resp.Status(http.StatusOK)

	// Try with the old password
	resp = he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	resp.JSON().Path("$.errors[0]").Equal("invalid email or password")

	// Try with the new password
	resp = he.POST(apiPath("/account/login")).
		WithJSON(model.LoginRequest{Email: testEmail, Password: newPassword}).
		Expect()
	resp.Status(http.StatusOK)
	token = resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)

}

func (s *AccountTestSuite) createUser(email string, password string) {
	stmts := []string{
		`INSERT INTO public.user_account(
	first_name, last_name, email, phone_number, active, updated_at, updated_by, version)
	VALUES ('Test', 'Test', $1, NULL, true, current_timestamp, 'test', 1) RETURNING ID`,
		`INSERT INTO public.user_credential(
	id, password_hash, expires_at, invalid_attempts, locked, activation_key, activation_key_expires_at, activated, 
reset_key, reset_key_expires_at, reset_at, updated_at, version)
	SELECT id, $1, $2, 0, false, null, null, true, NULL, NULL, NULL, current_timestamp, 1 FROM 
user_account where email = $3`,
	}

	mustExecStmt(s.DB, stmts[0], email)

	passwordHash, err := crypto.HashPassword(password)
	if err != nil {
		log.Printf("Error hasing password")
		panic(err)
	}
	mustExecStmt(s.DB, stmts[1], passwordHash, time.Now().Add(time.Second*1200), email)
}

func (s *AccountTestSuite) deleteUser(email string) {
	stmts := []string{
		`DELETE FROM user_credential WHERE id = (SELECT id FROM user_account where email = $1)`,
		`DELETE FROM user_account WHERE email = $1`,
	}

	mustExecStmt(s.DB, stmts[0], email)
	mustExecStmt(s.DB, stmts[1], email)

}
