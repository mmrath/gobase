package test

import (
	"bytes"
	"log"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"mmrath.com/gobase/pkg/model"
)

type AccountTestSuite struct {
	client_app.TestSuite
}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (s *AccountTestSuite) TestPing() {
	resp, err := http.Get(s.server.URL + "/ping")
	require.NoError(s.T(), err)
	require.Equal(s.T(), 200, resp.StatusCode)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	require.NoError(s.T(), err)
	contents := buf.String()
	require.Equal(s.T(), "pong", contents)
}

func (s *AccountTestSuite) TestSignUpActivateAndLogin() {
	he := httpexpect.New(s.T(), s.server.URL)
	testEmail := "test@example.com"
	testPassword := "Secret123"
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     testEmail,
		"password":  testPassword,
	}
	he.POST("/api/account/register").
		WithJSON(signupRequest).
		Expect().
		Status(http.StatusOK)

	// Login should throw an error now
	msg := s.mailer.PopLastMessage()
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Activate your account", msg.Subject)
	require.Equal(s.T(), signupRequest["email"], msg.To[0].Email)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: testPassword}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	resp.JSON().Path("$.details.cause").Equal("user is not activated")

	re := regexp.MustCompile("/account/activate\\?key=([0-9a-f\\-]+)")
	key := re.FindStringSubmatch(msg.Html)[1]
	he.GET("/api/account/activate").
		WithQuery("key", key).
		Expect().
		Status(http.StatusOK)

	resp = he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: testPassword}).
		Expect()
	resp.Status(http.StatusOK)
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]

	jwtService := pkg.NewJWTService(s.cfg.JWT)
	jwtToken, err := jwtService.Decode(token)
	require.NoError(s.T(), err)
	err = jwtToken.Claims.Valid()
	require.NoError(s.T(), err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	require.NotNil(s.T(), claims["jti"])
	require.Equal(s.T(), testEmail, claims["email"])
}

func (s *AccountTestSuite) TestSignUpWithInvalidEmail() {
	he := httpexpect.New(s.T(), s.server.URL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     "test234",
		"password":  "Secret123",
	}
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("email")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("must be a valid email address")
}

func (s *AccountTestSuite) TestSignUpWithDuplicateEmail() {
	he := httpexpect.New(s.T(), s.server.URL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     "test@example.com",
		"password":  "Secret123",
	}
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusOK)

	// 2nd Request
	resp = he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)

	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("email")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("user already exists")
}

func (s *AccountTestSuite) TestSignUpWithInvalidPassword() {
	he := httpexpect.New(s.T(), s.server.URL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     "test@example.com",
		"password":  "123", /// too short
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)

	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("the length must be between 6 and 32")
}

func (s *AccountTestSuite) TestSignUpWithLongPassword() {
	he := httpexpect.New(s.T(), s.server.URL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     "test@example.com",
		"password":  "12345678901234567890123456789012",
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusOK)
}

func (s *AccountTestSuite) TestSignUpWithTooLongPassword() {
	he := httpexpect.New(s.T(), s.server.URL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     "test@example.com",
		"password":  "12345678901234567890123456789012232",
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("the length must be between 6 and 32")
}

func (s *AccountTestSuite) TestActivateWithWrongKey() {
	he := httpexpect.New(s.T(), s.server.URL)
	resp := he.GET("/api/account/activate").
		WithQuery("key", "wrong-key").
		Expect().
		Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.cause").Equal("invalid activation token")
}

func (s *AccountTestSuite) TestResetPassword() {
	testEmail := "testuser@localhost"
	he := httpexpect.New(s.T(), s.server.URL)
	initResetRequest := map[string]interface{}{
		"email": testEmail,
	}
	resp := he.POST("/api/account/reset-password/init").
		WithJSON(initResetRequest).Expect()
	resp.Status(http.StatusOK)

	// Login should throw an error now
	msg := s.mailer.PopLastMessage()
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Reset password", msg.Subject)
	require.Equal(s.T(), initResetRequest["email"], msg.To[0].Email)

	re := regexp.MustCompile("/account/reset-password\\?key=([0-9a-f\\-]+)")
	key := re.FindStringSubmatch(msg.Html)[1]

	resetRequest := map[string]interface{}{
		"resetToken":  key,
		"newPassword": "Secret123",
	}

	he.POST("/api/account/reset-password/finish").
		WithJSON(resetRequest).
		Expect().
		Status(http.StatusOK)

	resp = he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: "Secret123"}).
		Expect()
	resp.Status(http.StatusOK)
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]

	jwtService := pkg.NewJWTService(s.cfg.JWT)
	jwtToken, err := jwtService.Decode(token)
	require.NoError(s.T(), err)
	err = jwtToken.Claims.Valid()
	require.NoError(s.T(), err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	require.NotNil(s.T(), claims["jti"])
	require.Equal(s.T(), testEmail, claims["email"])

}

func (s *AccountTestSuite) TestWithWrongUsername() {
	testEmail := "none@gobase.mmrath.com"
	he := httpexpect.New(s.T(), s.server.URL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: "Secret123"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.JSON().Path("$.details.cause").Equal("invalid email or password")
}

func (s *AccountTestSuite) TestWithWrongPassword() {
	testEmail := "none@gobase.mmrath.com"

	s.createUser(testEmail, "Secret123")

	he := httpexpect.New(s.T(), s.server.URL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: "incorrectPassword"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.JSON().Path("$.details.cause").Equal("invalid email or password")
}

func (s *AccountTestSuite) TestChangePassword() {
	testEmail := "testuser1@localhost"
	password := "Secret123"
	newPassword := "NewSecret123"
	s.createUser(testEmail, "Secret123")
	he := httpexpect.New(s.T(), s.server.URL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusOK)
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)

	resp = he.POST("/api/account/change-password").
		WithJSON(model.ChangePasswordRequest{CurrentPassword: password, NewPassword: newPassword}).
		WithHeader("Authorization", "Bearer "+token).
		Expect()
	resp.Status(http.StatusOK)

	// Try with the old password
	resp = he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	resp.JSON().Path("$.details.cause").Equal("invalid email or password")

	// Try with the new password
	resp = he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: newPassword}).
		Expect()
	resp.Status(http.StatusOK)
	token = resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)

}

func (s *AccountTestSuite) createUser(email string, password string) {
	stmts := []string{
		`INSERT INTO public.user_account(
	first_name, last_name, email, phone_number, active, created_at, created_by, updated_at, updated_by, version)
	VALUES ('Test', 'Test', ?, NULL, true, current_timestamp, 'test', current_timestamp, 'test', 1) RETURNING ID`,
		`INSERT INTO public.user_credential(
	id, password_hash, expires_at, invalid_attempts, locked, activation_key, activation_key_expires_at, activated, reset_key, reset_key_expires_at, reset_at, updated_at, version)
	SELECT id, ?, ?, 0, false, null, null, true, NULL, NULL, NULL, current_timestamp, 1 FROM user_account where email = ?`,
	}

	_, err := s.db.Exec(stmts[0], email)
	if err != nil {
		log.Printf("Error executing statement %s", stmts[0])
		panic(err)
	}
	passwordHash, err := pkg.HashPassword(pkg.SHA256([]byte(password)))
	if err != nil {
		log.Printf("Error hasing password")
		panic(err)
	}
	_, err = s.db.Exec(stmts[1], passwordHash, time.Now().Add(time.Second*1200), email)
	if err != nil {
		log.Printf("Error creating credentials")
		panic(err)
	}
}
