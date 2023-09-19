// Deadlock: https://yourbasic.org/golang/detect-deadlock/
package main

import "fmt"

func main() {

	// A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	fmt.Println(<-ch)

	ch2 := make(chan int)
	go func() {
		ch2 <- 1
	}()
	fmt.Println(<-ch2)
	// En goroutine kan fastna antingen för att den väntar på en kanal eller för att den väntar på ett av låsen i sync-paketet (mer om det nästa vecka)

	// Vanliga orsaker till deadlock är att ingen annan goroutine har tillgång till kanalen/låset eller att en grupp gorutiner väntar på varandra och ingen av dem kan fortsätta.

}
