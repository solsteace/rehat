package controllers

import (
	"net/http"
	"strconv"

	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/services"
)

type Admin struct {
	services.MotelManagement
	services.Auth
}

func (a Admin) Register(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	formData := req.PostForm
	newAdmin := models.User{
		Name:     formData.Get("name"),
		Username: formData.Get("username"),
		Password: []byte(formData.Get("password")),
		Email:    formData.Get("email"),
		Role:     "admin",
		IsActive: true}
	newAdmin, accessToken, err := a.Auth.Register(newAdmin)
	if err != nil {
		return err
	}

	payload := struct {
		models.User `json:"user"`
		AccessToken string `json:"authorization"`
	}{User: newAdmin, AccessToken: accessToken}
	if err := sendResponse(w, http.StatusCreated, payload); err != nil {
		return err
	}
	return nil
}

func (a Admin) AddMotel(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	userId, err := middlewares.TokenUserId(req.Context())
	if err != nil {
		return err
	}

	formData := req.PostForm
	motel := models.Motel{
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email"),
		Rating:        0}
	admin := models.MotelAdmin{UserID: userId}
	if err := a.MotelManagement.Add(&admin, &motel); err != nil {
		return err
	}

	payload := struct {
		Motel      models.Motel      `json:"motel"`
		MotelAdmin models.MotelAdmin `json:"motelAdmin"`
	}{Motel: motel, MotelAdmin: admin}
	if err := sendResponse(w, http.StatusCreated, payload); err != nil {
		return err
	}
	return nil
}

func (a Admin) EditMotelById(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}
	formData := req.PostForm

	userId, err := middlewares.TokenUserId(req.Context())
	if err != nil {
		return err
	}

	motelId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}

	motelRating, err := strconv.Atoi((formData.Get("rating")))
	if err != nil {
		return err
	}

	motel := models.Motel{
		MotelID:       motelId,
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email"),
		Rating:        motelRating}

	if err := a.MotelManagement.EditById(userId, motelId, &motel); err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motel); err != nil {
		return err
	}
	return nil
}

func (a Admin) DeleteMotelById(w http.ResponseWriter, req *http.Request) error {
	motelId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}

	userId, err := middlewares.TokenUserId(req.Context())
	if err != nil {
		return err
	}

	if err := a.MotelManagement.DeleteById(userId, motelId); err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusNoContent, nil); err != nil {
		return err
	}
	return nil
}
