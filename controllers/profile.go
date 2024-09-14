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
	userId, err := middlewares.TokenUserId(req.Context())
	if err != nil {
		return err
	}

	user, err := p.Service.Index(userId)
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusCreated, user); err != nil {
		return err
	}
	return nil
}
