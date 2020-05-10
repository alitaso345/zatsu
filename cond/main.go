package main

import (
	"fmt"
	"sync"
	"time"
)

var msg string = "hoge"

func main() {
	c := sync.NewCond(&sync.Mutex{})
	changed := false

	go func() {
		for {
			fmt.Printf("%s\n", msg)
		}
	}()
	go func() {
		messages := []string{"piyo", "huga", "nyan", "poyo"}
		for _, m := range messages {
			msg = m
			changed = true
			c.Signal()
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(20 * time.Second)
}
