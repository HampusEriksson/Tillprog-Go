// Gorutines
// https://yourbasic.org/golang/goroutines-explained/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Publish(text string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println(text)
}

func main() {
	// gorutine - thread - körs parallelt

	go Publish("Hello from a goroutine")
	Publish("Hello from main")

	go Publish("BREAKING NEWS: A goroutine starts a new thread.")
	fmt.Println("Let’s hope the news will published before I leave.")

	// Wait for the news to be published.
	fmt.Println("Program shutting down in 10 seconds...")
	time.Sleep(10 * time.Second)

	fmt.Println("Ten seconds has passed: I’m leaving now.")

}
