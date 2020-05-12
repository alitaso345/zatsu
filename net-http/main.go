package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":5000", nil)
}

type HelloJSON struct {
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "GET root path\n")
	case http.MethodPost:
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		var hello HelloJSON
		json.Unmarshal(buf.Bytes(), &hello)

		fmt.Printf("hello: %v\n", hello)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "POST root path\n")
	case http.MethodPut:
		w.WriteHeader(http.StatusNoContent)
		// 残りのメソッドもあるが省略
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed\n")
	}
}
