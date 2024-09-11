package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solsteace/rest/services"
)

func Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			token := req.Header.Get("authorization")
			if token == "" {
				payload := services.ErrAccessToken{
					Message: "No access token provided in header"}
				sendError(w, http.StatusUnauthorized, payload)
				return
			}

			const PREFIX = "Bearer "
			if token[:len(PREFIX)] != PREFIX {
				payload := services.ErrAccessToken{
					Message: fmt.Sprintf("Token should be prefixed with `%s`", PREFIX)}
				sendError(w, http.StatusBadRequest, payload)
				return
			}
			token = token[len(PREFIX):]
			req.Header.Set("authorization", token)

			parsedToken, err := jwt.ParseWithClaims(
				token,
				&services.TokenClaims{},
				func(t *jwt.Token) (interface{}, error) {
					// TODO: Ensure sign method is similar to what being used in server
					return []byte(os.Getenv("JWT_SECRET")), nil
				})
			if err != nil {
				payload := services.ErrAccessToken{Message: err.Error()}
				sendError(w, http.StatusBadRequest, payload)
				return
			}

			_, ok := parsedToken.Claims.(*services.TokenClaims)
			if !ok {
				payload := services.ErrAccessToken{Message: "No defined claim found within JWT"}
				sendError(w, http.StatusBadRequest, payload)
				return
			}

			next.ServeHTTP(w, req)
		},
	)
}
