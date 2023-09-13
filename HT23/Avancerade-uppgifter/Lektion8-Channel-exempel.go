package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create two channels for communication
	person1 := make(chan string)
	person2 := make(chan string)

	// Person 1 goroutine
	go chatPerson("Person 1", person1, person2)

	// Person 2 goroutine
	go chatPerson("Person 2", person2, person1)

	// Start the conversation
	person1 <- "Hello!"
	select {}
}

func chatPerson(name string, in, out chan string) {

	for {
		select {
		case msg := <-in:
			fmt.Printf("%s received: %s\n", name, msg)

			// Simulate a reply
			reply := getRandomWord()
			fmt.Printf("%s replying: %s\n", name, reply)
			out <- reply

			// Sleep to simulate typing delay
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))

		default:
			// Send a random word
			word := getRandomWord()
			fmt.Printf("%s sending: %s\n", name, word)
			out <- word

			// Sleep to simulate typing delay
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}
	}
}

func getRandomWord() string {
	words := []string{"Hello", "How", "Are", "You", "Good", "Morning", "Yes", "No", "Chat", "Program"}
	return words[rand.Intn(len(words))]
}
