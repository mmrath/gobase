package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
)

type userIdKeyType int

var userIdKey userIdKeyType

type JWTConfig struct {
	Secret         string        `mapstructure:"secret" yaml:"secret"`
	ExpiryDuration time.Duration `mapstructure:"expiryDuration" yaml:"expiryDuration"`
	CookieName string `mapstructure:"cookieName" yaml:"cookieName"`
	PubKeyPath string `mapstructure:"pubKeyPath" yaml:"pubKeyPath"`
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
	config    JWTConfig
	tokenAuth *jwtauth.JWTAuth
}

func NewJWTService(config JWTConfig) JWTService {
	return &jwtService{config: config, tokenAuth: jwtauth.New("HS256", []byte(config.Secret), nil)}
}
func (s *jwtService) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(s.tokenAuth)
}

func (s *jwtService) Decode(tokenString string) (t *jwt.Token, err error) {
	return s.tokenAuth.Decode(tokenString)
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
	exp := now.Add(s.config.ExpiryDuration).Unix()

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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Secret))
	return tokenString, err
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
