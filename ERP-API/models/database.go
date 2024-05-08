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

type Database struct {
	Conn *sql.DB
}

func (d *Database) Connect() {
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	d.Conn, err = sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = d.Conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func (d *Database) CloseConnection() {
	defer d.Conn.Close()
}
