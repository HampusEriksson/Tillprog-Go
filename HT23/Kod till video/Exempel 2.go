package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func simulateTask(tasks int, sum_chan chan int) {
	// Get a random number tasks times and add them to the sum

	sum := 0
	for i := 0; i < tasks; i++ {
		sum += rand.Intn(10)
	}
	sum_chan <- sum
}

func main() {
	startTime := time.Now()

	nr_of_tasks := int(math.Pow(2, 28)) // 268 435 456 - 268 million tasks

	nr_of_people := int(math.Pow(2, 7))

	sum_chan := make(chan int)

	for i := 0; i < nr_of_people; i++ {
		go simulateTask(nr_of_tasks/nr_of_people, sum_chan)
	}

	total_sum := 0
	for i := 0; i < nr_of_people; i++ {
		total_sum += <-sum_chan
	}
	fmt.Println("Sum:", total_sum)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Program execution time: %s\n", elapsedTime)
}
