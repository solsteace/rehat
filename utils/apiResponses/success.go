package apiResponses

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) error {
	body := response{Status: "Success", Data: data}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(payload)
	return nil
}
