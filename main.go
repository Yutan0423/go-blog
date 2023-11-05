package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	_, err0 := strconv.Atoi("a")
	fmt.Printf("err0: [%T] %v\n", err0, err0)
	err2 := strconv.ErrSyntax
	err3 := strconv.ErrRange
	fmt.Printf("err0: [%T] %v\n", err2, err2)
	fmt.Printf("err0: [%T] %v\n", err3, err3)

	fmt.Println(errors.Is(err0, err2))
	fmt.Println(errors.Is(err0, err3))

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
