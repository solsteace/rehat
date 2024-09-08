package controllers

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(w http.ResponseWriter, data interface{}) error {
	body := responseBody{Status: "Success", Data: data}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(payload)
	return nil
}
