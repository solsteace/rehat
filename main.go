package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	jwtLifetime, err := strconv.Atoi(os.Getenv("JWT_LIFETIME"))
	if err != nil {
		log.Fatalf("`JWT_LIFETIME` should be an integer")
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("Something went wrong during database setup: \n%s", err)
	}

	app := app{
		db: db,
		accessTokenCfg: struct {
			signMethod jwt.SigningMethod
			lifetime   time.Duration
			secret     string
		}{
			signMethod: jwt.SigningMethodHS256,
			lifetime:   time.Minute * time.Duration(jwtLifetime),
			secret:     os.Getenv("JWT_SECRET")}}
	app.init()
	server := http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT")),
		Handler: app.router,
	}

	log.Printf("Running server in port %s...", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Something went wrong when trying to run the server: \n%s", err)
	}
}

// Establishes database connection with given configuration in `.env`
func initDB() (*sql.DB, error) {
	var conn *sql.DB
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return conn, err
	}

	if err := conn.Ping(); err != nil {
		return conn, err
	}
	return conn, nil
}
