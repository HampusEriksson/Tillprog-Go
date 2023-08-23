package main

import (
	"fmt"
)

func main() {

	x := 5
	val := x < 5
	fmt.Printf("%t\n", val)

	// AND &&
	// OR ||
	// NOT !
	fmt.Println(x == 5 || x != 5 && x == 5)

	var user_age int
	fmt.Println("What's your age?")
	fmt.Scan(&user_age)

	// Samma jämförelser som Python. < > osv
	// < > <= >= != ==
	// Kan inte jämföra float och int
	// If
	// Else if
	// Else

	if user_age > 65 {
		fmt.Println("You are retired")
	} else if user_age >= 18 {
		fmt.Println("You are an adult")

	} else if user_age > 13 {
		fmt.Println("You are a youth")

	} else {
		fmt.Println("You are a baby")
	}
	// While - vad är det???
	// For
	// for x:=0; x<5; x++
	// break & continue - samma som python

	y := 5

	for y < 10 {
		y++
		fmt.Println(y)
	}

	for z := 0; z < 5; z++ {
		fmt.Println(z)
	}

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	// Switch
	// case x:
	// default:

	switch name {

	case "Jubin":
		fmt.Println("Hej")

	case "Leon":
		fmt.Println("Nej")

	case "Daniel":
		fmt.Println("Tåååg")

	default:
		fmt.Println("Inget case assignat")
	}

}
