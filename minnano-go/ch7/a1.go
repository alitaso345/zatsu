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

	row := db.QueryRow(`select name, age from users where id = $1`, 1)
	var name string
	var age int64
	err = row.Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(name, age)
}
