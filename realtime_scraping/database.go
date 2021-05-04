package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func storeDB() {
	connStr := "user=postgres password=cmpe295 dbname=panda host=localhost port=4000 sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}
