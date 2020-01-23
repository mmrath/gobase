package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"

	"github.com/mmrath/gobase/go/pkg/errutil"
)

type SsoClientConfig struct {
	CookieName string
	PubKeyPath string
}

type ssoMiddleware struct {
	PubKey     *rsa.PublicKey
	CookieName string
	URL        string
}

func NewSsoMiddleware(cfg SsoClientConfig) (func(handler http.Handler) http.Handler, error) {
	key, err := ioutil.ReadFile(cfg.PubKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load jwt public key")
		return nil, errutil.Wrap(err, "failed to load jwt public key file")
	}
	parsedPubKey, err := jwt.ParseRSAPublicKeyFromPEM(key)

	if err != nil {
		return nil, errutil.Wrap(err, "failed to parse public key from pem")
	}
	return ssoMiddleware{
		PubKey:     parsedPubKey,
		CookieName: cfg.CookieName,
	}.SsoMiddleware, nil
}

func (s *ssoMiddleware) SsoMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie(s.CookieName)
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "https://127.0.0.1:8081/sso?s_url=https://127.0.0.1:8082/cookie", http.StatusMovedPermanently)
			return
		}
		if err != nil {
			log.Error().Err(err).Msg("error reading cookie")
			errutil.RenderError(w, r, err)
			return
		}

		parts := strings.Split(strings.Split(c.String(), "=")[1], ".")
		err = jwt.SigningMethodRS512.Verify(strings.Join(parts[0:2], "."), parts[2], s.PubKey)
		if err != nil {
			log.Error().Err(err).Msg("error while verifying key")
			errutil.RenderError(w, r, err)
			return
		}

		tokenString := strings.Split(c.String(), "=")[1]
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return s.PubKey, nil
		})

		claims, ok := token.Claims.(*CustomClaims) // claims.User and claims.Roles are what we are interested in.
		if ok && token.Valid {
			fmt.Printf("User: %v Roles: %v Tok_Expires: %v \n", claims.Email, claims.Roles, claims.StandardClaims.ExpiresAt)

			next.ServeHTTP(w, r)

		} else {
			log.Error().Msg("invalid token")
			fmt.Println(err)
		}

	}
	return http.HandlerFunc(fn)
}

type CustomClaims struct {
	Email string   `json:"Address"`
	Roles []string `json:"Roles"`
	jwt.StandardClaims
}
