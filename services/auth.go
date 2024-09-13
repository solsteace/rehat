package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solsteace/rest/models"
	"golang.org/x/crypto/bcrypt"
)

type ErrAccessToken struct {
	Message string `json:"message"`
}

func (e ErrAccessToken) Error() string {
	return e.Message
}

type (
	AccessTokenCfg struct {
		SignMethod jwt.SigningMethod // TODO: Explore more about signing method
		Lifetime   time.Duration
		Secret     string
	}

	Auth struct {
		AccessTokenCfg
		Db *sql.DB
	}

	TokenClaims struct {
		UserId int64
		Role   string
		jwt.RegisteredClaims
	}
)

func (a Auth) LogIn(username, password string) (string, error) {
	var accessToken string

	u := models.User{}
	u, err := u.GetByUsername(a.Db, username)
	if err != nil {
		if _, ok := err.(*models.ErrRecordNotFound); !ok {
			return accessToken, err
		}
	}
	if u.IsNil() {
		return accessToken, errors.New("No user found with username `" + username + "`")
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return "", err
	}

	payload := TokenClaims{
		UserId: u.UserId,
		Role:   u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.AccessTokenCfg.Lifetime)),
		}}
	accessToken, err = a.generateJWT(a.AccessTokenCfg.Secret, payload)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (a Auth) Register(user models.User) (models.User, string, error) {
	var accessToken string

	u := models.User{}
	u, err := u.GetByUsername(a.Db, user.Username)
	if err != nil {
		if _, ok := err.(*models.ErrRecordNotFound); !ok {
			return user, accessToken, err
		}
	}
	if !u.IsNil() {
		return user, accessToken, &models.ErrDuplicateEntry{Field: "username"}
	}

	u, err = u.GetByEmail(a.Db, user.Email)
	if err != nil {
		if _, ok := err.(*models.ErrRecordNotFound); !ok {
			return user, accessToken, err
		}
	}
	if !u.IsNil() {
		return user, accessToken, &models.ErrDuplicateEntry{Field: "email"}
	}

	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, accessToken, err
	}

	user.Password = hash
	insertId, err := user.Save(a.Db)
	if err != nil {
		return user, accessToken, err
	}

	user.UserId = insertId
	user.Password = nil
	payload := TokenClaims{
		UserId: user.UserId,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.AccessTokenCfg.Lifetime)),
		}}
	accessToken, err = a.generateJWT(a.AccessTokenCfg.Secret, payload)
	if err != nil {
		return user, accessToken, err
	}

	return user, accessToken, nil
}

func (a Auth) generateJWT(secret string, payload TokenClaims) (string, error) {
	token := jwt.NewWithClaims(a.AccessTokenCfg.SignMethod, payload)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
