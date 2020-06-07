package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < 10; i++ {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
		time.Sleep(time.Second)
	}
	fmt.Println("finish main")
}
