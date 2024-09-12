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
	motel := controllers.Motel{Db: a.db}
	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", middlewares.HandleError(motel.GetById))
	motelApi.Handle("GET /", middlewares.HandleError(motel.GetAll))

	authService := services.Auth{Db: a.db, AccessTokenCfg: a.AccessTokenCfg}
	auth := controllers.Auth{Service: authService}
	authApi := http.NewServeMux()
	authApi.Handle("POST /login", middlewares.HandleError(auth.LogIn))
	authApi.Handle("POST /register", middlewares.HandleError(auth.Register))

	// TODO: Authorization based on resource ownership
	admin := controllers.Admin{Auth: authService}
	adminApi := http.NewServeMux()
	adminApi.Handle("POST /register", middlewares.HandleError(admin.Register))
	adminApi.Handle("POST /motel", middlewares.HandleError(admin.AddMotel))
	adminApi.Handle("PUT /motel/{id}", middlewares.HandleError(admin.EditMotelById))
	adminApi.Handle("DELETE /motel/{id}", middlewares.HandleError(admin.DeleteMotelById))

	api := http.NewServeMux()
	api.Handle("/motels/", http.StripPrefix("/motels", motelApi))
	api.Handle("/auth/", http.StripPrefix("/auth", authApi))
	api.Handle("/admin/", http.StripPrefix("/admin", adminApi))

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
