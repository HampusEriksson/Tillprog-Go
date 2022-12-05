// Channels : https://yourbasic.org/golang/channels-explained/
// Select: https://yourbasic.org/golang/select-explained/

package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel of ints
	//ic := make(chan int)

	// buffered channel with room for 10 strings
	//sc := make(chan string, 10)

	/*If the capacity of a channel is zero or absent,
	the channel is unbuffered and the sender blocks until the receiver has received the value.
	If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
	if the buffer is full, this means waiting until some receiver has retrieved a value.*/

	// A select statement blocks until at least one of it’s cases can proceed.

	ch := make(chan string, 1)
	go func() {
		ch <- "Hello!"
	}()

	ch2 := make(chan string, 1)
	go func() {
		ch2 <- "Hello!"
	}()

	// blocks until there's data available on ch1 or ch2
	select {
	case <-ch:
		fmt.Println("Received from ch1")
	case <-ch2:
		fmt.Println("Received from ch2")

	}

	//With zero cases this will never happen.
	// select {}

	// The default case is always able to proceed and runs if all other cases are blocked.
	select {
	case x := <-ch:
		fmt.Println("Received", x)
	default:
		fmt.Println("Nothing available")
	}

	select {
	case x := <-ch:
		fmt.Println("Received", x)
	case <-time.After(time.Second * 5):
		fmt.Println("Time out: No new info in 5 seconds")
	}
}
