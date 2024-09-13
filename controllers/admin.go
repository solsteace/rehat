package controllers

import (
	"database/sql"
	"net/http"

	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/services"
)

type Admin struct {
	Db *sql.DB
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
	newMotelId, err := motel.Save(a.Db)
	if err != nil {
		return err
	}

	motel.MotelID = newMotelId
	motelAdmin := models.MotelAdmin{
		UserID:  userId,
		MotelID: int64(motel.MotelID)}

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

// 	motelId, err := strconv.Atoi(req.PathValue("id"))
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
// 	_, err = motel.EditById(a.Db, req.PathValue("id"))
// 	if err != nil {
// 		return err
// 	}

// 	if err := sendResponse(w, http.StatusOK, motel); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a Admin) DeleteMotelById(w http.ResponseWriter, req *http.Request) error {
// 	motel := models.Motel{}
// 	err := motel.DeleteById(a.Db, req.PathValue("id"))
// 	if err != nil {
// 		return err
// 	}

// 	if err := sendResponse(w, http.StatusNoContent, nil); err != nil {
// 		return err
// 	}
// 	return nil
// }
