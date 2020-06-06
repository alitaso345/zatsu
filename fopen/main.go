package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fp := filepath.Join("./resources", os.Args[1])
	fs, err := os.Open(fp)
	if err != nil {
		log.Fatalln(err)
	}
	data := make([]byte, 100)
	count, err := fs.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
