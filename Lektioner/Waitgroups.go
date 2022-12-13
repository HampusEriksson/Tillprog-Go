package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go func() {
		// Do work.
		wg.Add(-1)
	}()
	go func() {
		// Do work.
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Waitgroup is done waiting")
}
