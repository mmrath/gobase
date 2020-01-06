package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mmrath/gobase/go/apps/uaa/internal/config"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	errorInvalidPassword = errors.New("invalid password")
	errorUserNotFound    = errors.New("user not found")
)

type Service struct {
	CookieName            string
	CookieDomain          string
	CookieValidityMinutes int64
	privateKey            *rsa.PrivateKey
}

// NewSSO creates a in-memory Auth provider
func NewSSO(cfg config.AuthConfig) (*Service, error) {
	privateKeyData, err := ioutil.ReadFile(cfg.JwtPrivateKeyPath)
	if err != nil {
		log.Err(err).Msg("failed to read jwt private key")
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Err(err).Msg("rsa key is invalid")
	}
	return &Service{
		CookieName:            cfg.CookieName,
		CookieDomain:          cfg.CookieDomain,
		CookieValidityMinutes: cfg.CookieValidityMinutes,
		privateKey:            privateKey,
	}, nil

}

func (m *Service) Auth(username string, password string) (User, []string, error) {
	if "foo@bar.com" == username && "password" == password {
		return User{UID: "123213", Email: username}, []string{"g1", "g2"}, nil
	}
	return User{}, nil, errors.New("bad username or password")
}

// BuildToken takes the user and the user roles info which is then signed by the private
// key of the login server. The expiry of the token is set per the third argument.
func (m *Service) BuildToken(user User, groups []string, _ time.Time) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(m.CookieValidityMinutes)).UTC()

	return genJwt(user, groups, m.privateKey, exp.Unix())
}

// BuildCookie takes the jwt token and returns a cookie and sets the expiration time of the same to that of
// the second arg.
func (m *Service) BuildCookie(s string, exp time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.CookieName,
		Value:    s,
		Domain:   m.CookieDomain,
		Path:     "/",
		Expires:  exp,
		MaxAge:   int(m.CookieValidityMinutes * 60),
		Secure:   true,
		HttpOnly: true,
	}
}

// Logout sets the expiration time of the cookie in the past rendering it unusable.
func (m *Service) Logout(expT time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.CookieName,
		Value:    "",
		Domain:   m.CookieDomain,
		Path:     "/",
		Expires:  expT,
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
}

func (m *Service) Is401(err error) bool {
	if err == errorInvalidPassword || err == errorUserNotFound {
		return true
	}
	return false
}
