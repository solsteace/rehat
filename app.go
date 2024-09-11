package main

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/solsteace/rest/controllers"
	"github.com/solsteace/rest/middlewares"
	"github.com/solsteace/rest/services"
)

type app struct {
	db     *sql.DB
	router *http.ServeMux
}

func (a *app) initDB() error {
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *app) initRoutes() error {
	motelService := services.Motel{Db: a.db}
	motel := controllers.Motel{Service: motelService}
	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", middlewares.HandleError(motel.GetById))
	motelApi.Handle("PUT /{id}", middlewares.HandleError(motel.Edit))
	motelApi.Handle("DELETE /{id}", middlewares.HandleError(motel.Delete))
	motelApi.Handle("GET /", middlewares.HandleError(motel.GetAll))
	motelApi.Handle("POST /", middlewares.HandleError(motel.Create))

	userService := services.User{Db: a.db}

	jwtLifetime, err := strconv.Atoi(os.Getenv("JWT_LIFETIME"))
	if err != nil {
		return err
	}

	authService := services.Auth{
		Db:   a.db,
		User: userService,
		AccessTokenCfg: services.AccessTokenCfg{
			SignMethod: jwt.SigningMethodHS256,
			Lifetime:   time.Minute * time.Duration(jwtLifetime),
			Secret:     os.Getenv("JWT_SECRET")}}
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
		http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			}))
	// router.Handle(
	// 	"/test",
	// 	middlewares.Log(
	// 		middlewares.Jwt(
	// 			http.HandlerFunc(
	// 				func(w http.ResponseWriter, req *http.Request) {
	// 					w.WriteHeader(http.StatusOK)
	// 					w.Write([]byte("YOU PASSED JWT TEST"))
	// 				}),
	// 		)))
	router.Handle(
		"/api/v1/",
		middlewares.Log(http.StripPrefix("/api/v1", api)))
	a.router = router

	return nil
}
