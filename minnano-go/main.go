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

	p1 := newPost("Go 1.1 released!", "Lorem ipsum lorem ipsum")
	p2 := newPost("Go 1.2 released!", "hoge piyo huga")

	err = dbmap.Insert(&p1, &p2)
	checkError(err, "Insert failed")

	// use convenience SelectInt
	count, err := dbmap.SelectInt("select count(*) from posts")
	checkError(err, "select count(*), failed")
	log.Println("Rows after inserting:", count)

	// update a row
	p2.Title = "Go 1.2 is better than ever"
	count, err = dbmap.Update(&p2)
	checkError(err, "Update failed")
	log.Println("Rows updated:", count)

	// fetch one row
	err = dbmap.SelectOne(&p2, "select * from posts where post_id=?", p2.Id)
	checkError(err, "SelectOne failed")
	log.Println("p2 row:", p2)

	// fetch all rows
	var posts []Post
	_, err = dbmap.Select(&posts, "select * from posts order by post_id")
	checkError(err, "select failed")
	log.Println("All rows:")
	for x, p := range posts {
		log.Printf(" %d: %v\n", x, p)
	}

	// delete row by PK
	count, err = dbmap.Delete(&p1)
	checkError(err, "Delete failed")
	log.Println("Rows deleted:", count)

	// delete row manually via Exec
	_, err = dbmap.Exec("delete from posts where post_id=?", p2.Id)
	checkError(err, "Exec failed")

	//confirm count is zero
	count, err = dbmap.SelectInt("select count(*) from posts")
	checkError(err, "select count(*) failed")
	log.Println("Row count - should be zero:", count)

	log.Println("Done!")
}

func newPost(title string, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

type Post struct {
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`
	Body    string `db:"article_body,size:1024"`
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./tmp/post_db.bin")
	checkError(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkError(err, "Create tables failed")

	return dbmap
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
