package main

import (
	"fmt"
	"sync"
	"time"
)

func parallel(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // mというMutexインスタンスで保護されたクリティカルセクションの専有を要求
	defer m.Unlock()

	fmt.Println("博")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("多")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("の")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("塩")
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go parallel(wg, m)
	}
	wg.Wait()
}
