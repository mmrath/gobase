package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	username string   `json:"username"`
	roles    []string `json:"roles"`
	jwt.StandardClaims
}

// genJwt generates the jwt token. Among other stuff, it packs in the authenticated user name and the roles that the
// user belongs to and an expiration time. The info is then signed by the private key of the login server.
func genJwt(u User, g []string, p *rsa.PrivateKey, t int64) (string, error) {
	claims := CustomClaims{
		u.Email,
		g,
		jwt.StandardClaims{
			ExpiresAt: t,
			Id:u.UID,
			Issuer:    "Login_Server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString(p)
}
