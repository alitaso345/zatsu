package main

import (
	"database/sql"
	"fmt"
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

	rows, err := db.Query(`select id, name, age from users order by age desc`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int64
		var name string
		var age int64
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, age)
	}
}
