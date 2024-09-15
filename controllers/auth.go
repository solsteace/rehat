package controllers

import (
	"net/http"

	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/services"
)

type Auth struct {
	Service services.Auth
}

func (a Auth) LogIn(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	form := req.PostForm
	accessToken, err := a.Service.LogIn(form.Get("username"), form.Get("password"))
	if err != nil {
		return err
	}

	payload := struct {
		AccessToken string `json:"authorization"`
	}{AccessToken: accessToken}
	if err := sendResponse(w, http.StatusOK, payload); err != nil {
		return err
	}
	return nil
}

func (a Auth) Register(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	formData := req.PostForm
	newUser := models.User{
		Name:     formData.Get("name"),
		Username: formData.Get("username"),
		Password: []byte(formData.Get("password")),
		Email:    formData.Get("email"),
		Role:     "customer"}
	newUser, accessToken, err := a.Service.Register(newUser)
	if err != nil {
		return err
	}

	payload := struct {
		models.User `json:"user"`
		AccessToken string `json:"authorization"`
	}{User: newUser, AccessToken: accessToken}
	if err := sendResponse(w, http.StatusCreated, payload); err != nil {
		return err
	}
	return nil
}
