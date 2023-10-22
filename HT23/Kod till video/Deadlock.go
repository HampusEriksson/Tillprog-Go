// Deadlock: https://yourbasic.org/golang/detect-deadlock/
package main

import "fmt"

func main() {

	// En goroutine kan fastna  för att den väntar på en kanal
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	fmt.Println(<-ch)

	ch2 := make(chan int)
	go func() {
		ch2 <- 1
	}()
	fmt.Println(<-ch2)

}
