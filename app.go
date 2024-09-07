package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
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

func (a *app) initRoutes() {
	motel := controllers.Motel{Service: services.Motel{Db: a.db}}
	motelApi := http.NewServeMux()
	motelApi.Handle("GET /{id}", middlewares.HandleError(motel.GetById))
	motelApi.Handle("PUT /{id}", middlewares.HandleError(motel.Edit))
	motelApi.Handle("DELETE /{id}", middlewares.HandleError(motel.Delete))
	motelApi.Handle("GET /", middlewares.HandleError(motel.GetAll))
	motelApi.Handle("POST /", middlewares.HandleError(motel.Create))

	api := http.NewServeMux()
	api.Handle("/motels/", http.StripPrefix("/motels", motelApi))

	router := http.NewServeMux()
	router.Handle(
		"/health",
		http.HandlerFunc(
			func(w http.ResponseWriter, req *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			}))
	router.Handle(
		"/api/v1/",
		middlewares.Log(http.StripPrefix("/api/v1", api)))
	a.router = router
}
