// Data races: https://yourbasic.org/golang/data-races-explained/
// Detect data races: https://yourbasic.org/golang/detect-data-races/

package main

import (
	"fmt"
)

func main() {
	//race()
	fmt.Println(getNumber())
	sharingIsCaring()
}

func sharingIsCaring() {
	ch := make(chan int)
	n := 0
	go func() {
		// A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n = <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
}

func getNumber() int {
	// https://www.sohamkamani.com/golang/data-races/
	i := 0
	go func() {
		i = 5
	}()

	return i
}
