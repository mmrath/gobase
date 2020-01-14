package tests

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"github.com/mmrath/gobase/go/pkg/crypto"
	"log"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/dgrijalva/jwt-go"
	"github.com/gavv/httpexpect/v2"
	"github.com/mmrath/gobase/go/pkg/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	TestSuite
}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (s *AccountTestSuite) TestPing() {
	resp, err := http.Get(s.ClipoURL + "/ping")
	require.NoError(s.T(), err)
	require.Equal(s.T(), 200, resp.StatusCode)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	require.NoError(s.T(), err)
	contents := buf.String()
	require.Equal(s.T(), "pong", contents)
}

func (s *AccountTestSuite) TestSignUpActivateAndLogin() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	testEmail := gofakeit.Email()
	testPassword := gofakeit.Password(true, true, true, true, true, 8)
	registerRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     testEmail,
		"password":  testPassword,
	}
	fmt.Printf("registering with data %v\n", registerRequest)
	he.POST("/api/account/register").
		WithJSON(registerRequest).
		Expect().
		Status(http.StatusOK)

	msg := s.EmailClient.GetLatestEmail(testEmail)
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Activate your account", msg.Subject)
	require.Equal(s.T(), registerRequest["email"], msg.To[0].Email)

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

	var jwtPublicKey *rsa.PublicKey
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		return jwtPublicKey, nil
	})
	require.NoError(s.T(), err)
	err = jwtToken.Claims.Valid()
	require.NoError(s.T(), err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	require.NotNil(s.T(), claims["jti"])
	require.Equal(s.T(), testEmail, claims["sub"])
}

func (s *AccountTestSuite) TestSignUpWithInvalidEmail() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	signupRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     "invalid_email",
		"password":  gofakeit.Password(true, true, true, true, true, 8),
	}
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("email")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("must be a valid email address")
}

func (s *AccountTestSuite) TestSignUpWithDuplicateEmail() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	signupRequest := map[string]interface{}{
		"firstName": gofakeit.FirstName(),
		"lastName":  gofakeit.LastName(),
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true,true,true,true,true,8),
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
	he := httpexpect.New(s.T(), s.ClipoURL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true,true,true,true,false,5), /// too short
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)

	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("the length must be between 6 and 32")
}

func (s *AccountTestSuite) TestSignUpWithLongPassword() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true,true,true,true,true,20),
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusOK)
}

func (s *AccountTestSuite) TestSignUpWithTooLongPassword() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	signupRequest := map[string]interface{}{
		"firstName": "Murali",
		"lastName":  "Rath",
		"email":     gofakeit.Email(),
		"password":  gofakeit.Password(true,true,true,true,true,32),
	}
	// 2nd Request
	resp := he.POST("/api/account/register").WithJSON(signupRequest).Expect()
	resp.Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.fieldErrors[0].field").Equal("password")
	resp.JSON().Path("$.details.fieldErrors[0].message").Equal("the length must be between 6 and 32")
}

func (s *AccountTestSuite) TestActivateWithWrongKey() {
	he := httpexpect.New(s.T(), s.ClipoURL)
	resp := he.GET("/api/account/activate").
		WithQuery("key", "wrong-key").
		Expect().
		Status(http.StatusBadRequest)
	resp.JSON().Path("$.details.cause").Equal("invalid activation token")
}

func (s *AccountTestSuite) TestResetPassword() {
	testEmail := gofakeit.Email()
	he := httpexpect.New(s.T(), s.ClipoURL)
	initResetRequest := map[string]interface{}{
		"email": testEmail,
	}
	resp := he.POST("/api/account/reset-password/init").
		WithJSON(initResetRequest).Expect()
	resp.Status(http.StatusOK)

	// Login should throw an error now
	msg := s.EmailClient.GetLatestEmail(testEmail)
	require.NotNil(s.T(), msg)
	require.Equal(s.T(), "Reset password", msg.Subject)
	require.Equal(s.T(), initResetRequest["email"], msg.To[0])

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

	var jwtPublicKey *rsa.PublicKey
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
		return jwtPublicKey, nil
	})
	require.NoError(s.T(), err)
	err = jwtToken.Claims.Valid()
	require.NoError(s.T(), err)
	claims := jwtToken.Claims.(jwt.MapClaims)
	require.NotNil(s.T(), claims["jti"])
	require.Equal(s.T(), testEmail, claims["sub"])

}

func (s *AccountTestSuite) TestWithWrongUsername() {
	testEmail := gofakeit.Email()
	he := httpexpect.New(s.T(), s.ClipoURL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: "Secret123"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.JSON().Path("$.details.cause").Equal("invalid email or password")
}

func (s *AccountTestSuite) TestWithWrongPassword() {
	testEmail := gofakeit.Email()
	password := gofakeit.Password(true,true,true,true,true,8)
	s.createUser(testEmail, password)
	defer s.deleteUser(testEmail)

	he := httpexpect.New(s.T(), s.ClipoURL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: "incorrectPassword"}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	_ = resp.Header("Authorization").NotMatch("Bearer (.*)")
	resp.Text().Contains("invalid email or password")
}

func (s *AccountTestSuite) TestChangePassword() {
	testEmail := gofakeit.Email()
	password := gofakeit.Password(true,true,true,true,true,8)
	newPassword := gofakeit.Password(true,true,true,true,true,10)
	s.createUser(testEmail, password)
	defer s.deleteUser(testEmail)

	he := httpexpect.New(s.T(), s.ClipoURL)

	resp := he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusOK)

	fmt.Printf("Cookies %v", resp.Cookies().Iter())
	jwtCookie := resp.Cookie("jwt").Value().Raw()
	token := resp.Header("Authorization").Match("Bearer (.*)").Raw()[1]
	require.NotNil(s.T(), token)

	resp = he.POST("/api/account/change-password").
		WithJSON(model.ChangePasswordRequest{CurrentPassword: password, NewPassword: newPassword}).
		WithHeader("Authorization", "Bearer "+token).
		WithCookie("jwt", jwtCookie).
		Expect()
	resp.Status(http.StatusOK)

	// Try with the old password
	resp = he.POST("/api/account/login").
		WithJSON(model.LoginRequest{Email: testEmail, Password: password}).
		Expect()
	resp.Status(http.StatusUnauthorized)
	resp.Text().Contains("invalid email or password")

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
	first_name, last_name, email, phone_number, active, updated_at, updated_by, version)
	VALUES ('Test', 'Test', $1, NULL, true, current_timestamp, 'test', 1) RETURNING ID`,
		`INSERT INTO public.user_credential(
	id, password_hash, expires_at, invalid_attempts, locked, activation_key, activation_key_expires_at, activated, reset_key, reset_key_expires_at, reset_at, updated_at, version)
	SELECT id, $1, $2, 0, false, null, null, true, NULL, NULL, NULL, current_timestamp, 1 FROM user_account where email = $3`,
	}

	mustExecStmt(s.db, stmts[0], email)
	passwordHash, err := crypto.HashPassword(crypto.SHA256([]byte(password)))
	if err != nil {
		log.Printf("Error hasing password")
		panic(err)
	}
	mustExecStmt(s.db, stmts[1], passwordHash, time.Now().Add(time.Second*1200), email)
}

func (s *AccountTestSuite) deleteUser(email string) {
	stmts := []string{
		`DELETE FROM user_credential WHERE id = (SELECT id FROM user_account where email = $1)`,
		`DELETE FROM user_account WHERE email = $1`,
	}

	mustExecStmt(s.db, stmts[0], email)
	mustExecStmt(s.db, stmts[1], email)

}
