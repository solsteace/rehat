package controllers

import (
	"net/http"

	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/services"
)

type Profile struct {
	Service services.Profile
}

func (p Profile) Index(w http.ResponseWriter, req *http.Request) error {
	userInfo, err := middlewares.UserFromToken(req.Context())
	if err != nil {
		return err
	}

	user, err := p.Service.Index(userInfo.Id)
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusCreated, user); err != nil {
		return err
	}
	return nil
}
