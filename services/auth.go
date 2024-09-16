package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
	"golang.org/x/crypto/bcrypt"
)

type (
	Auth struct {
		AccessToken
		repositories.User
	}

	UserInfo struct {
		Id   int64
		Role string
	}
)

func (a Auth) LogIn(username, password string) (string, error) {
	var accessToken string

	user, err := a.User.GetByUsername(username)
	if err != nil {
		return accessToken, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return accessToken, err
	}

	payload := AccessTokenClaims{
		UserInfo: UserInfo{
			Id:   user.UserId,
			Role: user.Role,
		}}
	accessToken, err = a.AccessToken.Generate(payload)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (a Auth) Register(user models.User) (models.User, string, error) {
	var accessToken string

	_, err := a.User.GetByUsername(user.Username)
	if err != nil {
		if _, ok := err.(*repositories.ErrRecordNotFound); !ok {
			return user, accessToken, err
		}
	} else {
		return user, accessToken, &repositories.ErrDuplicateEntry{Field: "username"}
	}

	_, err = a.User.GetByEmail(user.Email)
	if err != nil {
		if _, ok := err.(*repositories.ErrRecordNotFound); !ok {
			return user, accessToken, err
		}
	} else {
		return user, accessToken, &repositories.ErrDuplicateEntry{Field: "email"}
	}

	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return user, accessToken, err
	}

	user.Password = hash
	insertId, err := a.User.Save(user)
	if err != nil {
		return user, accessToken, err
	}

	user.UserId = insertId
	user.Password = nil
	payload := AccessTokenClaims{
		UserInfo: UserInfo{
			Id:   user.UserId,
			Role: user.Role,
		}}
	accessToken, err = a.AccessToken.Generate(payload)
	if err != nil {
		return user, accessToken, err
	}

	return user, accessToken, nil
}
