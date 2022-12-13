// Data races: https://yourbasic.org/golang/data-races-explained/
// Detect data races: https://yourbasic.org/golang/detect-data-races/

package main

import (
	"fmt"
)

func main() {
	// go run -race Lektion9del2.go
	//race()
	fmt.Println(getNumber())
}

func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait
	fmt.Println(n) // Output: <unspecified>
}

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
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
