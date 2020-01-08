package auth

import (
	"context"
	"crypto/rsa"
	"github.com/mmrath/gobase/go/pkg/errutil"
	"io/ioutil"
	"net/http"
	"strings"
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
	TokenValidityDuration time.Duration `default:"7d" split_words:"true"`
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

	privateKeyData, err := ioutil.ReadFile(config.PrivateKeyPath)
	if err != nil {
		return nil, errutil.Wrap(err, "failed to read jwt private key")
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, errutil.Wrap(err, "jwt rsa private key is invalid")
	}

	publicKeyData, err := ioutil.ReadFile(config.PublicKeyPath)
	if err != nil {
		return nil, errutil.Wrap(err, "failed to read jwt public key")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		return nil, errutil.Wrap(err, "jwt rsa public  key is invalid")
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
	return token,nil
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
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if token == nil || !token.Valid {
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


func fromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil // No error, just no token
	}

	// TODO: Make this a bit more robust, parsing-wise
	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errutil.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}