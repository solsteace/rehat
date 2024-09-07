package middlewares

import (
	"log"
	"net/http"

	"github.com/solsteace/rest/utils/responses"
)

type routeHandler func(w http.ResponseWriter, req *http.Request) error

func HandleError(handler routeHandler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			err := handler(w, req)
			if err == nil {
				return
			}

			// Determine error type here

			err = responses.Failure(w, 500, struct {
				Message string `json:"message"`
			}{Message: err.Error()})

			if err != nil {
				log.Printf("%s\n", err.Error())
				w.WriteHeader(500)
				w.Write([]byte("Internal server error"))
			}
		},
	)
}
