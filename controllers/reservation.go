package controllers

import (
	"net/http"

	"github.com/solsteace/rest/services"
)

type Reservation struct {
	Service services.Reservation
}

func (p Reservation) Add(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (p Reservation) EditById(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (p Reservation) DeleteById(w http.ResponseWriter, req *http.Request) error {
	return nil
}
