package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, _ := r.FormFile("file")
	buf, _ := ioutil.ReadAll(file)
	ioutil.WriteFile(fmt.Sprintf("./files/%s", fileHeader.Filename), buf, 0600)
}

func main() {
	http.HandleFunc("/upload", upload)
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
