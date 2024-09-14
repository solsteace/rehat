package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/solsteace/rest/repositories"
	"github.com/solsteace/rest/services"
)

type routeHandler func(w http.ResponseWriter, req *http.Request) error

// Handles error that returned by the route handler, if there's any.
// The error received would be analyzed to determine appropiate response to be
// sent to the client
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

// Produces more user-friendly error message (Might be useful in production)
func getProductionMessage(e error) string {
	switch e.(type) {
	// TODO: Complete
	}
	return ""
}

// Defines appropiate status code given the error received
func getErrStatusCode(e error) int {
	switch e.(type) {
	// === 400 BadRequest
	case *repositories.ErrDuplicateEntry:
		return http.StatusBadRequest

	// === 401 Unauthorized
	// case *services.ErrDuplicateEntry:

	// === 404 NotFound
	case *repositories.ErrRecordNotFound:
		return http.StatusNotFound

	// === 501 NotImplemented
	case *services.ErrNotImplemented:
		return http.StatusNotImplemented

	// === 500 InternalServerError (kept for clarity)
	case *repositories.ErrSQL:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
