// Lektion 7 - Förkunskaper till Uppgift 1 - 2
// Channels
// https://yourbasic.org/golang/channels-explained/

package main

import (
	"fmt"
)

func main() {

	// Skapar en kanal för strängar
	ch := make(chan string)

	// Skriver ut data från kanalen med en gorutine
	go func() {
		ch <- "Hello!"

		close(ch)
	}()

	fmt.Println(<-ch) // Print "Hello!".
	fmt.Println(<-ch) // Print the zero value "" without blocking.

	// Buffered channel
	int_ch := make(chan int, 10)

	int_ch <- 5

	fmt.Println(<-int_ch)

	ch2 := make(chan string)

	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			ch2 <- input

			if input == "exit" {
				break
			}
		}

		close(ch2)
	}()

	// Receive values from ch until closed.
	for value := range ch2 {
		fmt.Println("Från kanalen:", value) // Will not be executed.
	}

}
