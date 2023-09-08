// Lektion 7 - Förkunskaper till Uppgift 1 - 2
// Channels
// https://yourbasic.org/golang/channels-explained/

package main

import (
	"fmt"
)

func main() {
	/*If the capacity of a channel is zero or absent,
	    the channel is unbuffered and the sender blocks until the receiver has received the value.
	  If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
	  if the buffer is full, this means waiting until some receiver has retrieved a value.
	  Receivers always block until there is data to receive.*/

	//int_ch := make(chan int)
	//string_ch := make(chan string, 10)

	ch := make(chan string)
	go func() {
		ch <- "Hello!"

		close(ch)
	}()

	fmt.Println(<-ch) // Print "Hello!".
	fmt.Println(<-ch) // Print the zero value "" without blocking.
	fmt.Println(<-ch) // Once again print "".
	//v, ok := <-ch     // v is "", ok is false because channel is closed.

	// Receive values from ch until closed.
	for v := range ch {
		fmt.Println("Från kanalen:", v) // Will not be executed.
	}
}
