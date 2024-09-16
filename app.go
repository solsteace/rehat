package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/solsteace/rest/controllers"
	mw "github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/repositories"
	"github.com/solsteace/rest/services"
)

type app struct {
	db             *sql.DB
	router         *http.ServeMux
	accessTokenCfg struct {
		signMethod jwt.SigningMethod
		lifetime   time.Duration
		secret     string
	}
}

// Initializes apps by setting up routes, handlers, services, and more.
func (a *app) init() {
	userRepo := repositories.User{Db: a.db}
	motelRepo := repositories.Motel{Db: a.db}
	motelAdminRepo := repositories.MotelAdmin{Db: a.db}
	reservationRepo := repositories.Reservation{Db: a.db}
	roomRepo := repositories.Room{Db: a.db}
	classRepo := repositories.Class{Db: a.db}

	accessTokenService := services.AccessToken{
		SignMethod: a.accessTokenCfg.signMethod,
		Lifetime:   a.accessTokenCfg.lifetime,
		Secret:     a.accessTokenCfg.secret}
	authService := services.Auth{
		AccessToken: accessTokenService,
		User:        userRepo}
	profileService := services.Profile{UserRepo: userRepo}
	motelManagementService := services.MotelManagement{
		Motel:      motelRepo,
		MotelAdmin: motelAdminRepo}
	reservationService := services.Reservation{
		Room:        roomRepo,
		Reservation: reservationRepo,
		Class:       classRepo}

	motel := controllers.Motel{MotelRepo: motelRepo}
	auth := controllers.Auth{Service: authService}
	admin := controllers.Admin{
		Auth:            authService,
		MotelManagement: motelManagementService}
	profile := controllers.Profile{Service: profileService}
	reservation := controllers.Reservation{Service: reservationService}

	// TIP: Trace routing from the bottom then work your way up

	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", mw.HandleError(motel.GetById))
	motelApi.Handle("GET /", mw.HandleError(motel.GetAll))

	authApi := http.NewServeMux()
	authApi.Handle("POST /login", mw.HandleError(auth.LogIn))
	authApi.Handle("POST /register", mw.HandleError(auth.Register))

	adminMotelApi := http.NewServeMux()
	adminMotelApi.Handle("PUT /{id}", mw.HandleError(admin.EditMotel))
	adminMotelApi.Handle("DELETE /{id}", mw.HandleError(admin.DeleteMotel))
	adminMotelApi.Handle("POST /", mw.HandleError(admin.AddMotel))

	adminApi := http.NewServeMux()
	adminApi.Handle("POST /register", mw.HandleError(admin.Register))
	adminApi.Handle(
		"/motels/",
		http.StripPrefix("/motels", mw.Jwt(mw.Admin(adminMotelApi))))

	profileApi := http.NewServeMux()
	profileApi.Handle("/", mw.HandleError(profile.Index))

	reservationApi := http.NewServeMux()
	reservationApi.Handle("PUT /{id}", mw.HandleError(reservation.EditById))
	reservationApi.Handle("DELETE /{id}", mw.HandleError(reservation.DeleteById))
	reservationApi.Handle("POST /", mw.HandleError(reservation.Add))

	api := http.NewServeMux()
	api.Handle("/motels/", http.StripPrefix("/motels", motelApi))
	api.Handle("/auth/", http.StripPrefix("/auth", authApi))
	api.Handle("/admin/", http.StripPrefix("/admin", adminApi))
	api.Handle("/profile/", http.StripPrefix("/profile", mw.Jwt(profileApi)))
	api.Handle("/reservation/", http.StripPrefix("/reservation", mw.Jwt(reservationApi)))

	router := http.NewServeMux()
	router.Handle(
		"/health",
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}))
	router.Handle(
		"/api/v1/",
		mw.Log(http.StripPrefix("/api/v1", api)))

	a.router = router
}
