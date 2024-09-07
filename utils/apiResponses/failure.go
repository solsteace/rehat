package apiResponses

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Failure(w http.ResponseWriter, statusCode int, reason interface{}) error {
	if http.StatusText(statusCode) == "" {
		message := fmt.Sprintf("Status code unknown: %d", statusCode)
		return errors.New(message)
	}

	body := response{Status: "Failed", Data: reason}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode) // TODO: Bad request thingy
	w.Write(payload)
	return nil
}
