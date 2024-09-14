package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solsteace/rest/services"
)

type (
	middlewareKey string
	tokenContent  struct {
		userId   int64
		userRole string
	}
)

const tokenKey middlewareKey = "token"

// Checks for JWT in the `authorization` header. If it's not present, an error
// message would be sent.
//
// Upon checking, if the content is appropiate, it would be decoded and stored
// in requests' context. An error message would be sent when the decoding process
// is failed
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

			claims, ok := parsedToken.Claims.(*services.TokenClaims)
			if !ok {
				payload := services.ErrAccessToken{Message: "No defined claim found within JWT"}
				sendError(w, http.StatusBadRequest, payload)
				return
			}

			ctx := context.WithValue(
				context.Background(),
				middlewareKey(tokenKey),
				&tokenContent{
					userId:   claims.UserId,
					userRole: claims.Role})
			next.ServeHTTP(w, req.WithContext(ctx))
		},
	)
}

func TokenUserId(ctx context.Context) (int64, error) {
	v, ok := ctx.Value(tokenKey).(*tokenContent)
	if !ok {
		return 0, errors.New("no `userId` found")
	}
	return v.userId, nil
}

func TokenUserRole(ctx context.Context) (string, error) {
	v, ok := ctx.Value(tokenKey).(*tokenContent)
	if !ok {
		return "", errors.New("no `userRole` found")
	}
	return v.userRole, nil
}
