package services

import (
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
)

type Profile struct {
	UserRepo repositories.User
}

func (p Profile) Index(userId int64) (models.User, error) {
	var user models.User
	user, err := p.UserRepo.GetById(userId)
	if err != nil {
		return user, err
	}
	return user, nil
}
