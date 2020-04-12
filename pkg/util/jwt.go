package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired.")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type JWT interface {
	Generate(claims *Claims) (string, error)
	Parse(token string) (*Claims, error)
	Refresh(token string) (string, error)
}

func NewJWT() JWT {
	return &Token{
		[]byte("dreamland"),
	}
}

type Token struct {
	Secret []byte
}

type Claims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Generate 生成token
func (t *Token) Generate(claims *Claims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		Issuer:    string(t.Secret),
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(t.Secret)
}

// Parse 解析token
func (t *Token) Parse(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})
	if err != nil {
		if value, ok := err.(*jwt.ValidationError); ok {
			switch {
			case value.Errors&jwt.ValidationErrorMalformed != 0:
				err = TokenMalformed
			case value.Errors&jwt.ValidationErrorExpired != 0:
				err = TokenExpired
			case value.Errors&jwt.ValidationErrorNotValidYet != 0:
				err = TokenNotValidYet
			default:
				err = TokenInvalid
			}
			return nil, err
		}
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// Refresh 刷新token
func (t *Token) Refresh(token string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return t.Generate(claims)
	}
	return "", TokenInvalid
}
