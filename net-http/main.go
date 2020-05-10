package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":5000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "GET root path\n")
	case http.MethodPost:
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
