package controllers

import (
	"net/http"
	"strconv"

	"github.com/solsteace/rest/repositories"
)

type Motel struct {
	MotelRepo repositories.Motel
}

func (m Motel) GetAll(w http.ResponseWriter, req *http.Request) error {
	motels, err := m.MotelRepo.GetAll()
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motels); err != nil {
		return err
	}
	return nil
}

func (m Motel) GetById(w http.ResponseWriter, req *http.Request) error {
	motelId, err := strconv.ParseInt(req.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}

	motel, err := m.MotelRepo.GetById(motelId)
	if err != nil {
		return err
	}

	if err := sendResponse(w, http.StatusOK, motel); err != nil {
		return err
	}
	return nil
}
