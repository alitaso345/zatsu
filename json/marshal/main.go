package main

import (
	"encoding/json"
	"os"
)

type Category struct {
	Id   int    `json:"id"`
	Name string `json:name`
}

type Book struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}

func main() {
	book := Book{
		Id:   1,
		Name: "alitaso book",
		Categories: []Category{
			{Id: 3, Name: "programing"},
			{Id: 4, Name: "tech"},
		},
	}

	json.NewEncoder(os.Stdout).Encode(&book)
}
