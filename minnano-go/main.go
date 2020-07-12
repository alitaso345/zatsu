package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v2"
)

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkError(err, "TruncateTables failed")

	// create new users
	user1 := newUser("alitaso346")
	user2 := newUser("xin")
	user3 := newUser("beatrooper")
	user4 := newUser("niya")

	// insert rows
	err = dbmap.Insert(&user1, &user2, &user3, &user4)
	checkError(err, "Insert failed")

	// update a row
	user1.Name = "きこーだ"
	user4.Name = "DJ NIi-ya"
	_, err = dbmap.Update(&user1, &user4)
	checkError(err, "Update failed")

	// fetch all rows
	var users []User
	_, err = dbmap.Select(&users, "select * from users order by user_id")
	checkError(err, "Select failed")
	log.Println("All artists...")
	for i, u := range users {
		log.Printf("%d: %s\n", i, u.Name)
	}

	log.Println("Done!")
}

func newUser(name string) User {
	return User{
		Created: time.Now().UnixNano(),
		Name:    name,
	}
}

type User struct {
	Id      int64 `db:"user_id"`
	Created int64
	Name    string `db:",size:50"`
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./tmp/user_db.bin")
	checkError(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkError(err, "Create tables failed")

	return dbmap
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
