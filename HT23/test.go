package main

import "fmt"

func main() {
	numbers := make([]int, 5, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 1)
	fmt.Println(numbers)
}
