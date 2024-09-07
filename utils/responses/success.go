package responses

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) error {
	body := body{Status: "Success", Data: data}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(200)
	w.Write(payload)
	return nil
}
