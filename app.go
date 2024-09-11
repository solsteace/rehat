package main

import (
	"database/sql"
	"net/http"

	"github.com/solsteace/rest/controllers"
	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/services"
)

type app struct {
	db     *sql.DB
	router *http.ServeMux
	services.AccessTokenCfg
}

func (a *app) init() {
	motelService := services.Motel{Db: a.db}
	motel := controllers.Motel{Service: motelService}
	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", middlewares.HandleError(motel.GetById))
	motelApi.Handle("PUT /{id}", middlewares.HandleError(motel.Edit))
	motelApi.Handle("DELETE /{id}", middlewares.HandleError(motel.Delete))
	motelApi.Handle("GET /", middlewares.HandleError(motel.GetAll))
	motelApi.Handle("POST /", middlewares.HandleError(motel.Create))

	userService := services.User{Db: a.db}

	authService := services.Auth{
		Db:             a.db,
		User:           userService,
		AccessTokenCfg: a.AccessTokenCfg}
	auth := controllers.Auth{Service: authService}
	authApi := http.NewServeMux()
	authApi.Handle("POST /login", middlewares.HandleError(auth.LogIn))
	authApi.Handle("POST /register", middlewares.HandleError(auth.Register))

	api := http.NewServeMux()
	api.Handle("/motels/", http.StripPrefix("/motels", motelApi))
	api.Handle("/auth/", http.StripPrefix("/auth", authApi))

	router := http.NewServeMux()
	router.Handle(
		"/health",
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}))
	router.Handle(
		"/api/v1/",
		middlewares.Log(http.StripPrefix("/api/v1", api)))
	a.router = router
}
