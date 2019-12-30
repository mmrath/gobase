package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/mmrath/gobase/go/pkg/auth"
)

// genJwt generates the jwt token. Among other stuff, it packs in the authenticated user name and the roles that the
// user belongs to and an expiration time. The info is then signed by the private key of the login server.
func genJwt(u User, permissions []string, p *rsa.PrivateKey, t int64) (string, error) {
	claims := auth.CustomClaims{
		Email: u.Email,
		Roles: permissions,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t,
			Id:        u.UID,
			Issuer:    "Login_Server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString(p)
}
