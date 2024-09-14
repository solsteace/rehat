package middlewares

import (
	"log"
	"net/http"
)

// Logs the request received to the console. The log would be printed
// in <Method> | <URL> format
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			log.Printf("%s | %s\n", req.Method, req.URL)
			next.ServeHTTP(w, req)
		},
	)
}
