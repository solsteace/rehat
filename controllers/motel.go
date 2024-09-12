package controllers

import (
	"database/sql"
	"net/http"

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
