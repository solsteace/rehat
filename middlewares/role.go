package middlewares

import (
	"net/http"
)

// Checks whether the user account has `Admin` role. Successful check will allow
// the user to proceed to the next handler, while Failed check would send code 403
// response instead.
//
// `Superadmin` role bypasses this check
func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			userInfo, err := UserFromToken(req.Context())
			if err != nil {
				sendError(w, http.StatusBadRequest, struct {
					Message string `json:"message"`
				}{Message: err.Error()})
				return
			}

			role := userInfo.Role
			if role != "admin" && role != "superadmin" {
				sendError(w, http.StatusForbidden, struct {
					Message string `json:"message"`
				}{Message: "Not enough role previlege"})
				return
			}
			next.ServeHTTP(w, req)
		},
	)
}

// Checks whether the user account has `Superadmin` role. Successful check will allow
// the user to proceed to the next handler, while Failed check would send code 403
// response instead
func Superadmin(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			userInfo, err := UserFromToken(req.Context())
			if err != nil {
				sendError(w, http.StatusBadRequest, struct {
					Message string `json:"message"`
				}{Message: err.Error()})
				return
			}

			role := userInfo.Role
			if role != "superadmin" {
				sendError(w, http.StatusForbidden, struct {
					Message string `json:"message"`
				}{Message: "Not enough role previlege"})
				return
			}
			next.ServeHTTP(w, req)
		},
	)
}
