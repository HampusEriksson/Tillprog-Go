package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	// For i in range(1000)
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {

	wg.Add(3)
	go increment()
	go increment()
	go increment()
	wg.Wait()

	fmt.Println("Counter:", counter)
}
