package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	twiceBody := func() {
		fmt.Println("twice")
	}
	// 異なる値を渡されたとしても、同一のインスタンスを使っている限り初回の呼び出しのみが実行される
	// つまり以下の行を追加しても、文字列は表示されない
	once.Do(twiceBody)
}
