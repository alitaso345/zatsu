package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	_ = http.ListenAndServe(":8080", nil)
}