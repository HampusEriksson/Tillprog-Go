// Lektion 7 - Förkunskaper till Uppgift 1 - 2
// Gorutines
// https://yourbasic.org/golang/goroutines-explained/

package main

import (
	"fmt"
	"time"
)

func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
	}()
}

func main() {
	// thread - körs parallelt
	// gorutine - lightweight thread

	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")
	fmt.Println("Hello from main goroutine")
	fmt.Println("Hello from main goroutine")
	time.Sleep(time.Second * 2)

	Publish("A goroutine starts a new thread.", 5*time.Second)
	fmt.Println("Let’s hope the news will published before I leave.")

	// Wait for the news to be published.
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds later: I’m leaving now.")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Hour(), ":", time.Now().Minute())

}
