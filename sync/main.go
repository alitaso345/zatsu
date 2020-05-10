package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var m sync.Map

	go func() {
		for {
			m.Store(rand.Intn(100000), struct{}{})
		}
	}()

	for {
		fmt.Println(m.Load(rand.Intn(100000)))
	}
}
