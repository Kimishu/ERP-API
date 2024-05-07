package database

import (
	"database/sql"
	"fmt"
	_ "gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "cyb"
	dbname   = "ERP-db"
)

func Connect() *sql.DB {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
	return db
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
