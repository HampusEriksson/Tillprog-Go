package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup

func increment(ch chan struct{}, done chan bool) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
	ch <- struct{}{}
	done <- true
}

func main() {
	ch := make(chan struct{})
	done := make(chan bool)
	wg.Add(2)
	go increment(ch, done)
	go increment(ch, done)
	wg.Wait()
	close(ch)
	<-done // Wait for both goroutines to finish
	<-done
	fmt.Println("Counter:", counter)
}
