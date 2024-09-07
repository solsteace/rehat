package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/solsteace/rest/services"
	"github.com/solsteace/rest/utils/apiResponses"
)

type Motel struct {
	Service services.Motel
}

func (r Motel) GetAll(w http.ResponseWriter, req *http.Request) error {
	motels, err := r.Service.GetAll()
	if err != nil {
		return err
	}

	if err := apiResponses.Success(w, motels); err != nil {
		return err
	}
	return nil
}

func (r Motel) GetById(w http.ResponseWriter, req *http.Request) error {
	motel, err := r.Service.GetById(req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := apiResponses.Success(w, motel); err != nil {
		return err
	}
	return nil
}

func (r Motel) Create(w http.ResponseWriter, req *http.Request) error {
	motel, err := r.Service.Create()
	if err != nil {
		return err
	}

	data, err := json.Marshal(motel)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(data)
	return nil
}

func (r Motel) Edit(w http.ResponseWriter, req *http.Request) error {
	motel, err := r.Service.EditById(req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := apiResponses.Success(w, motel); err != nil {
		return err
	}
	return nil
}

func (r Motel) Delete(w http.ResponseWriter, req *http.Request) error {
	err := r.Service.DeleteById(req.PathValue("id"))
	if err != nil {
		return err
	}

	if err := apiResponses.Success(w, nil); err != nil {
		return err
	}
	return nil
}
