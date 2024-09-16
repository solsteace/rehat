package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	AccessToken struct {
		SignMethod jwt.SigningMethod // TODO: Explore more about signing method
		Lifetime   time.Duration
		Secret     string
	}

	AccessTokenClaims struct {
		UserInfo
		jwt.RegisteredClaims
	}

	ErrAccessToken struct {
		Message string `json:"message"`
	}
)

func (e ErrAccessToken) Error() string {
	return e.Message
}

func (a AccessToken) Generate(payload AccessTokenClaims) (string, error) {
	payload.RegisteredClaims = jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.Lifetime)),
	}

	token := jwt.NewWithClaims(a.SignMethod, payload)
	tokenString, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
