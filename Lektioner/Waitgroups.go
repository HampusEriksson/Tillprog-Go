package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// Do work.
		wg.Add(2)
	}()
	go func() {
		// Do work.
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Waitgroup is done waiting")
}
