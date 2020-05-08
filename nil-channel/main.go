package main

import (
	"fmt"
	"time"
)

func streamingAPI(ch chan<- string, quit chan<- bool) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("data: %d", i)
		fmt.Println("address: ", &msg)
		ch <- msg
		time.Sleep(1 * time.Second)
	}

	quit <- true
}

func main() {
	quit := make(chan bool)

	// 外部のAPIからデータを取得してきて、チャネルで送信する

	// 取得してきたデータを、以下にあるgoroutineに同じく送信したい
	// 利用イメージとしてはhttpHandleFunc
	go func() {
		ch := make(chan string)
		go streamingAPI(ch, quit)
		for {
			select {
			case msg := <-ch:
				fmt.Println("go routine address: ", &msg)
				fmt.Printf("HandlerA: %s\n", msg)
			}
		}
	}()
	//go func() {
	//	ch := make(chan string)
	//	go streamingAPI(ch, quit)
	//	for {
	//		select {
	//		case msg := <-ch:
	//			fmt.Printf("HandlerB: %s\n", msg)
	//		}
	//	}
	//}()

	<-quit
}

// 上記スクリプトの出力結果
//HandlerB: data: 0
//HandlerA: data: 1
//HandlerB: data: 2
//HandlerA: data: 3
//HandlerB: data: 4

//以下のようになるように修正したい
//HandlerB: data: 0
//HandlerA: data: 0
//HandlerB: data: 1
//HandlerA: data: 1
//HandlerB: data: 2
//HandlerA: data: 2
//HandlerB: data: 3
//HandlerA: data: 3
//HandlerB: data: 4
//HandlerA: data: 4
