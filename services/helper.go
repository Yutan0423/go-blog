package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func init() {
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	fmt.Println(dbConn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
