package controllers

import (
	"net/http"
	"strconv"

	"github.com/solsteace/rest/models"
	"github.com/solsteace/rest/services"
)

type Motel struct {
	Service services.Motel
}

func (m Motel) GetAll(w http.ResponseWriter, req *http.Request) error {
	motels, err := m.Service.GetAll()
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motels); err != nil {
		return err
	}
	return nil
}

func (m Motel) GetById(w http.ResponseWriter, req *http.Request) error {
	motel, err := m.Service.GetById(req.PathValue("id"))
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
	newMotel := models.Motel{
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email")}
	newMotelId, err := m.Service.Create(newMotel)
	if err != nil {
		return err
	}

	newMotel.MotelID = int(newMotelId)
	if err := sendResponse(w, http.StatusCreated, newMotel); err != nil {
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
	newMotel := models.Motel{
		MotelID:       motelId,
		Name:          formData.Get("name"),
		Location:      formData.Get("location"),
		ContactNumber: formData.Get("contactNumber"),
		Email:         formData.Get("email")}
	_, err = m.Service.EditById(req.PathValue("id"), newMotel)
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, newMotel); err != nil {
		return err
	}
	return nil
}

func (m Motel) Delete(w http.ResponseWriter, req *http.Request) error {
	err := m.Service.DeleteById(req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusNoContent, nil); err != nil {
		return err
	}
	return nil
}
