package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Yutan0423/go-medium-level/api"
)

var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Println("failed to connect db")
		return
	}
	r := api.NewRouter(db)

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func init() {
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
