package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/solsteace/rest/services"
)

type routeHandler func(w http.ResponseWriter, req *http.Request) error

func HandleError(handler routeHandler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			err := handler(w, req)
			if err == nil {
				return
			}

			statusCode := getErrStatusCode(err)
			message := err.Error()
			if os.Getenv("ENVIRONMENT") != "DEVEL" {
				message = getProductionMessage(err)
			}

			data := struct {
				Message string `json:"message"`
			}{Message: message}
			err = sendError(w, statusCode, data)
			if err != nil {
				log.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
			}
		},
	)
}

// Might be useful in production for more user-friendly response
func getProductionMessage(e error) string {
	switch e.(type) {
	// TODO: Complete
	}
	return ""
}

func getErrStatusCode(e error) int {
	statusCode := http.StatusInternalServerError
	switch e.(type) {
	// === 400 BadRequest
	case *services.ErrDuplicateEntry:
		statusCode = http.StatusBadRequest

	// === 401 Unauthorized
	// case *services.ErrDuplicateEntry:

	// === 501 NotImplemented
	case *services.ErrServiceNotImplemented:
		statusCode = http.StatusNotImplemented

	// === 500 InternalServerError
	case *services.ErrSQL: // Skip or keep for clarity?
		statusCode = http.StatusInternalServerError
	}
	return statusCode
}
