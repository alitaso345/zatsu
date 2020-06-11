package main

import (
	"fmt"
	"log"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

func main() {
	e := env.NewEnv()

	err := e.Define("println", fmt.Println)
	if err != nil {
		log.Fatalf("Define error: %v\n", err)
	}

	script := `
println("Hello World :)")
`
	_, err = vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("Execute error: %v\n", err)
	}
}
