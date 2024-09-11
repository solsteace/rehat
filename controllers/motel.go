package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/solsteace/rest/models"
)

type Motel struct {
	Db *sql.DB
}

func (m Motel) GetAll(w http.ResponseWriter, req *http.Request) error {
	motel := models.Motel{}
	motels, err := motel.GetAll(m.Db)
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motels); err != nil {
		return err
	}
	return nil
}

func (m Motel) GetById(w http.ResponseWriter, req *http.Request) error {
	motel := models.Motel{}
	motel, err := motel.GetById(m.Db, req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motel); err != nil {
		return err
	}
	return nil
}

func (m Motel) Create(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	formData := req.PostForm
	motel := models.Motel{
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email")}
	newMotelId, err := motel.Save(m.Db)
	if err != nil {
		return err
	}

	motel.MotelID = int(newMotelId)
	if err := sendResponse(w, http.StatusCreated, motel); err != nil {
		return err
	}
	return nil
}

func (m Motel) Edit(w http.ResponseWriter, req *http.Request) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	motelId, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		return err
	}

	formData := req.PostForm
	motel := models.Motel{
		MotelID:       motelId,
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email")}
	_, err = motel.EditById(m.Db, req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motel); err != nil {
		return err
	}
	return nil
}

func (m Motel) Delete(w http.ResponseWriter, req *http.Request) error {
	motel := models.Motel{}
	err := motel.DeleteById(m.Db, req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusNoContent, nil); err != nil {
		return err
	}
	return nil
}
