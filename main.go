package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := app{}
	if err := app.initDB(); err != nil {
		log.Fatalf("Something went wrong during database setup: \n%s", err)
	}

	app.initRoutes()
	server := http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT")),
		Handler: app.router,
	}

	log.Printf("Running server in port %s...", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Something went wrong during server setup: \n%s", err)
	}
}
