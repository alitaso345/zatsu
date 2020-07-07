package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=127.0.0.1 port=5432 user=test password=password dbname=test sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(`insert into users(name, age) values('Bob', 18)`)
	if err != nil {
		log.Fatal(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID: %v", affected)
}
