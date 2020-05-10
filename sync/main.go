package main

import (
	"fmt"
	"sync"
	"time"
)

var userMap sync.Map

func streamingAPI(quit chan<- bool) {
	for i := 0; i < 5; i++ {
		userMap.Range(func(key, val interface{}) bool {
			ch, ok := val.(chan string)
			if ok {
				ch <- fmt.Sprintf("data: %d", i)
			} else {
				fmt.Printf("why error for %v\n", key)
			}
			return true
		})
		time.Sleep(1 * time.Second)
	}
	quit <- true
}

func addUser(user string) {
	ch := make(chan string)
	userMap.Store(user, ch)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Printf("Handler %s: %s\n", user, msg)
			}
		}
	}()
}

func main() {
	quit := make(chan bool)
	go streamingAPI(quit)

	addUser("foo")

	time.Sleep(1 * time.Second)
	addUser("bar")

	time.Sleep(1 * time.Second)
	addUser("baz")

	time.Sleep(1 * time.Second)
	addUser("quux")
	<-quit
}
