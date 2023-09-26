package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup
var mu sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go increment()
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}
