package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mmrath/gobase/pkg/error_util"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
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

func NewSsoMiddleware(cfg SsoClientConfig) (*ssoMiddleware, error) {
	key, err := ioutil.ReadFile(cfg.PubKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load jwt public key")
		return nil, error_util.NewInternal(err, "failed to load jwt public key file")
	}
	parsedPubKey, err := jwt.ParseRSAPublicKeyFromPEM(key)

	return &ssoMiddleware{
		PubKey:     parsedPubKey,
		CookieName: cfg.CookieName,
	}, nil
}

func (s *ssoMiddleware) SsoMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie(s.CookieName)
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "https://127.0.0.1:8081/sso?s_url=https://127.0.0.1:8082/cookie", 301)
			return
		}
		if err != nil {
			log.Error().Err(err).Msg("error reading cookie")
			error_util.RenderError(w, r, err)
			return
		}

		parts := strings.Split(strings.Split(c.String(), "=")[1], ".")
		err = jwt.SigningMethodRS512.Verify(strings.Join(parts[0:2], "."), parts[2], s.PubKey)
		if err != nil {
			log.Error().Err(err).Msg("error while verifying key")
			error_util.RenderError(w, r, err)
			return
		}

		tokenString := strings.Split(c.String(), "=")[1]
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return s.PubKey, nil
		})

		claims, ok := token.Claims.(*CustomClaims) // claims.User and claims.Roles are what we are interested in.
		if ok && token.Valid {
			fmt.Printf("User: %v Roles: %v Tok_Expires: %v \n", claims.Email, claims.Roles, claims.StandardClaims.ExpiresAt)

			next.ServeHTTP(w,r)

		} else {
			log.Error().Msg("invalid token")
			fmt.Println(err)
		}

	}
	return http.HandlerFunc(fn)
}

type CustomClaims struct {
	Email string   `json:"Email"`
	Roles []string `json:"Roles"`
	jwt.StandardClaims
}
