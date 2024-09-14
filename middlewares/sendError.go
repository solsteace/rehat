package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type payload struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Sends error message to the client with specified `statusCode`. `reason` could
// be used to provide information to the client regarding the reason of failure
func sendError(w http.ResponseWriter, statusCode int, reason interface{}) error {
	if http.StatusText(statusCode) == "" {
		message := fmt.Sprintf("Status code unknown: %d", statusCode)
		return errors.New(message)
	}

	body := payload{Status: "Failed", Data: reason}
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	w.Write(payload)
	return nil
}
