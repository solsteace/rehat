package controllers

import (
	"net/http"

	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/repositories"
	"github.com/solsteace/rest/services"
)

type Admin struct {
	MotelRepo      repositories.Motel
	MotelAdminRepo repositories.MotelAdmin
	UserRepo       repositories.User
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

	// TODO: Make transaction
	userId, err := middlewares.TokenUserId(req.Context())
	if err != nil {
		return err
	}

	formData := req.PostForm
	motel := models.Motel{
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email")}
	newMotelId, err := a.MotelRepo.Save(motel)
	if err != nil {
		return err
	}

	motel.MotelID = newMotelId
	motelAdmin := models.MotelAdmin{
		UserID:  userId,
		MotelID: int64(motel.MotelID)}
	newMotelAdminId, err := a.MotelAdminRepo.Save(motelAdmin)
	if err != nil {
		return err
	}

	motelAdmin.AdminID = newMotelAdminId
	payload := struct {
		Motel      models.Motel      `json:"motel"`
		MotelAdmin models.MotelAdmin `json:"motelAdmin"`
	}{Motel: motel, MotelAdmin: motelAdmin}
	if err := sendResponse(w, http.StatusCreated, payload); err != nil {
		return err
	}
	return nil
}

// func (a Admin) EditMotelById(w http.ResponseWriter, req *http.Request) error {
// 	err := req.ParseForm()
// 	if err != nil {
// 		return err
// 	}

// 	motelId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	formData := req.PostForm
// 	motel := models.Motel{
// 		MotelID:       motelId,
// 		Name:          formData.Get("name"),
// 		Location:      formData.Get("location"),
// 		ContactNumber: formData.Get("contactNumber"),
// 		Email:         formData.Get("email")}
// 	_, err = a.Motel.EditById(motelId, motel)
// 	if err != nil {
// 		return err
// 	}

// 	if err := sendResponse(w, http.StatusOK, motel); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a Admin) DeleteMotelById(w http.ResponseWriter, req *http.Request) error {
// 	motelId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}

// 	if err := a.Motel.DeleteById(motelId); err != nil {
// 		return err
// 	}

// 	if err := sendResponse(w, http.StatusNoContent, nil); err != nil {
// 		return err
// 	}
// 	return nil
// }
