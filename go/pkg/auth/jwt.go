package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type userIdKeyType int

var userIdKey userIdKeyType

type JWTConfig struct {
	CookieName            string        `default:"jwt" split_words:"true"`
	CookieDomain          string        `split_words:"true"`
	TokenValidityDuration time.Duration `default:"240h" split_words:"true"`
	PrivateKeyPath        string        `split_words:"true"`
	PublicKeyPath         string        `split_words:"true"`
}

type JWTService interface {
	Verifier() func(http.Handler) http.Handler
	NewToken(user Principal) (tokenString string, err error)
	Decode(tokenString string) (t *jwt.Token, err error)
	Authenticator(http.Handler) http.Handler
}

type Principal interface {
	GetName() string
	GetEmail() string
	GetId() int64
}

type jwtService struct {
	cookieName            string
	cookieDomain          string
	tokenValidityDuration time.Duration
	privateKey            *rsa.PrivateKey
	publicKey             *rsa.PublicKey
	jwtAuth               *jwtauth.JWTAuth
}

func NewJWTService(config JWTConfig) (JWTService, error) {

	var privateKey *rsa.PrivateKey
	var publicKey *rsa.PublicKey

	if config.PublicKeyPath != "" && config.PrivateKeyPath != "" {
		privateKeyData, err := ioutil.ReadFile(config.PrivateKeyPath)
		if err != nil {
			return nil, errutil.Wrap(err, "failed to read jwt private key")
		}
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
		if err != nil {
			return nil, errutil.Wrap(err, "jwt rsa private key is invalid")
		}

		publicKeyData, err := ioutil.ReadFile(config.PublicKeyPath)
		if err != nil {
			return nil, errutil.Wrap(err, "failed to read jwt public key")
		}

		publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
		if err != nil {
			return nil, errutil.Wrap(err, "jwt rsa public  key is invalid")
		}
	} else {
		key, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return nil, errutil.Wrap(err, "failed to generate key pair")
		}
		privateKey = key
		publicKey = &key.PublicKey
	}

	return &jwtService{
		cookieName:            config.CookieName,
		cookieDomain:          config.CookieDomain,
		tokenValidityDuration: config.TokenValidityDuration,
		privateKey:            privateKey,
		publicKey:             publicKey,
		jwtAuth:               jwtauth.New("RS512", privateKey, publicKey)}, nil
}
func (s *jwtService) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(s.jwtAuth)
}

func (s *jwtService) Decode(tokenString string) (t *jwt.Token, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errutil.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})

	if err != nil {
		return nil, errutil.Wrap(err, "failed to parse jwt token")
	}
	return token, nil
}

func UserIdFromContext(ctx context.Context) int64 {
	id := ctx.Value(userIdKey).(int64)
	return id
}

func NewAuthContext(ctx context.Context, userId int64) context.Context {
	ctx = context.WithValue(ctx, userIdKey, userId)
	return ctx
}

type Claims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func (s *jwtService) NewToken(user Principal) (string, error) {
	now := time.Now()
	iat := now.Unix()
	exp := now.Add(s.tokenValidityDuration).Unix()

	claims := &Claims{
		UserId: user.GetId(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: exp,
			IssuedAt:  iat,
			Id:        uuid.New().String(),
			Subject:   user.GetEmail(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", errutil.Wrap(err, "failed to sign jwt token")
	}
	return tokenString, nil
}

func (s *jwtService) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			log.Error().Err(err).Send()
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if token == nil || !token.Valid {
			log.Error().Msg("token is not valid")
			http.Error(w, http.StatusText(401), 401)
			return
		}

		userId, ok := claims["userId"]

		if !ok {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		req := r.WithContext(NewAuthContext(r.Context(), int64(userId.(float64))))

		// Token is authenticated, pass it through
		next.ServeHTTP(w, req)
	})
}
