package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int64
	Name MyNullString
	Age  int64
}

type MyNullString struct {
	s sql.NullString
}

func (s *MyNullString) Scan(value interface{}) error {
	return s.s.Scan(value)
}

func (s MyNullString) String() string {
	if !s.s.Valid {
		return "nil"
	}
	return s.s.String
}

func (s MyNullString) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.s.String)
}

func (s *MyNullString) UnmarshalJSON(data []byte) error {
	var ss string
	if err := json.Unmarshal(data, &ss); err != nil {
		s.s.String = ""
		s.s.Valid = false
		return err
	}
	s.s.String = ss
	s.s.Valid = true
	return nil
}

func main() {
	dsn := "host=127.0.0.1 port=5432 user=test password=password dbname=test sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`select id, name, age from users`)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(&user)
		fmt.Println(user)
	}
}
