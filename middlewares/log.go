package middlewares

import (
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			log.Printf("%s | %s\n", req.Method, req.URL)
			next.ServeHTTP(w, req)
		},
	)
}
