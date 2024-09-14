package main

import (
	"database/sql"
	"net/http"

	"github.com/solsteace/rest/controllers"
	mw "github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/repositories"
	"github.com/solsteace/rest/services"
)

type app struct {
	db     *sql.DB
	router *http.ServeMux
	services.AccessTokenCfg
}

func (a *app) init() {
	userRepo := repositories.User{Db: a.db}
	motelRepo := repositories.Motel{Db: a.db}
	motelAdminRepo := repositories.MotelAdmin{Db: a.db}

	authService := services.Auth{
		User:           userRepo,
		AccessTokenCfg: a.AccessTokenCfg}
	profileService := services.Profile{UserRepo: userRepo}

	motel := controllers.Motel{MotelRepo: motelRepo}
	auth := controllers.Auth{Service: authService}
	admin := controllers.Admin{
		MotelRepo:      motelRepo,
		MotelAdminRepo: motelAdminRepo,
		UserRepo:       userRepo,
		Auth:           authService}
	profile := controllers.Profile{Service: profileService}

	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", mw.HandleError(motel.GetById))
	motelApi.Handle("GET /", mw.HandleError(motel.GetAll))

	authApi := http.NewServeMux()
	authApi.Handle("POST /login", mw.HandleError(auth.LogIn))
	authApi.Handle("POST /register", mw.HandleError(auth.Register))

	adminMotelApi := http.NewServeMux()
	adminMotelApi.Handle("POST /", mw.HandleError(admin.AddMotel))
	// adminMotelApi.Handle("PUT /{id}", mw.HandleError(admin.EditMotelById))
	// adminMotelApi.Handle("DELETE /{id}", mw.HandleError(admin.DeleteMotelById))

	adminApi := http.NewServeMux()
	adminApi.Handle("POST /register", mw.HandleError(admin.Register))
	adminApi.Handle(
		"POST /motels/",
		http.StripPrefix("/motels", mw.Jwt(mw.Admin(adminMotelApi))))

	profileApi := http.NewServeMux()
	profileApi.Handle("/", mw.HandleError(profile.Index))

	api := http.NewServeMux()
	api.Handle("/motels/", http.StripPrefix("/motels", motelApi))
	api.Handle("/auth/", http.StripPrefix("/auth", authApi))
	api.Handle("/admin/", http.StripPrefix("/admin", adminApi))
	api.Handle("/profile/", http.StripPrefix("/profile", mw.Jwt(profileApi)))

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
