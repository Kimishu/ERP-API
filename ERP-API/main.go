package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "dbname=ERP-db user=postgres password=??? sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, z := db.Query("SELECT name FROM \"Enterprises\" WHERE subscription_id = (SELECT id FROM \"Subscriptions\" WHERE name = 'standard')")
	if z != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name string
		)
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name: %s\n", name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
