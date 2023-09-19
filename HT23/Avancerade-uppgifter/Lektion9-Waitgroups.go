package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(4)
	go func() {
		// Do work.
		wg.Add(-1)
	}()
	go func() {
		// Do work.
		wg.Done()
	}()
	go do_work()
	go do_work()

	wg.Wait()
	fmt.Println("Waitgroup is done waiting")
}

func do_work() {
	// Do work.
	wg.Done()
}
