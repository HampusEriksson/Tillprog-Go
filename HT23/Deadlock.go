// Deadlock: https://yourbasic.org/golang/detect-deadlock/
package main

import "fmt"

func main() {

	// A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)

	// A goroutine can get stuck
	// either because it’s waiting for a channel or
	// because it is waiting for one of the locks in the sync package.

	// Common reasons are that
	// no other goroutine has access to the channel or the lock
	// a group of goroutines are waiting for each other and none of them is able to proceed.

}
