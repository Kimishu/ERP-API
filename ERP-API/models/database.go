package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "cyb"
	dbname   = "ERP-db"
)

var Database *sql.DB

func ConnectDB() {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	Database, err = sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = Database.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
