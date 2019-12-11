package auth

import (
	"crypto/rsa"
	"errors"
	"github.com/mmrath/gobase/apps/sso/pkg/config"
	"net/http"
	"time"
)

var(
	errorInvalidPassword= errors.New("invalid password")
	errorUserNotFound = errors.New("user not found")
)

type SSOProvider struct {
	Cookie *config.CookieConfig
	User   string
	Pass   string
	Groups *[]string
	privateKey *rsa.PrivateKey
}

// NewMemorySSO creates a in-memory SSO provider
func NewMemorySSO(cfg *config.CookieConfig) (*SSOProvider, error) {

	return &SSOProvider{
		Cookie: cfg,
		User:   "alice",
		Pass:   "password123",
		Groups: &[]string{"admin", "moderator", "super"},
	}, nil

}

// Auth takes user,password strings as arguments and returns the user, user roles (e.g ldap groups)
// (string slice) if the call succeeds. Auth should return the ErrUnAuthorized or ErrUserNotFound error if
// auth fails or if the user is not found respectively.
func (m *SSOProvider) Auth(username string, password string) (User, []string, error) {
	if "foo@bar.com" == username && "password" == password {
		return User{UID: "123213", Email: username}, []string{"g1", "g2"}, nil
	}
	return User{}, nil, errors.New("bad username or password")
}

// TokenValidityMinutes returns the cookie/jwt token validity in hours.
func (m *SSOProvider) TokenValidityMinutes() int64 {
	return m.Cookie.ValidHours
}

func (m *SSOProvider) CookieName() string {
	return m.Cookie.Name
}

func (m *SSOProvider) CookieDomain() string {
	return m.Cookie.Domain
}

// BuildToken takes the user and the user roles info which is then signed by the private
// key of the login server. The expiry of the token is set per the third argument.
func (m *SSOProvider) BuildToken(user User, groups []string, _ time.Time) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(m.TokenValidityMinutes())).UTC()

	return genJwt(user, groups, m.privateKey, exp.Unix())
}

// BuildCookie takes the jwt token and returns a cookie and sets the expiration time of the same to that of
// the second arg.
func (m *SSOProvider) BuildCookie(s string, exp time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.Cookie.Name,
		Value:    s,
		Domain:   m.Cookie.Domain,
		Path:     "/",
		Expires:  exp,
		MaxAge:   int(m.Cookie.ValidHours * 3600),
		Secure:   true,
		HttpOnly: true,
	}
}

// Logout sets the expiration time of the cookie in the past rendering it unusable.
func (m *SSOProvider) Logout(expT time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.Cookie.Name,
		Value:    "",
		Domain:   m.Cookie.Domain,
		Path:     "/",
		Expires:  expT,
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
}

func (m *SSOProvider) Is401(err error) bool{
	if err == errorInvalidPassword || err == errorUserNotFound {
		return true
	}
	return false
}

