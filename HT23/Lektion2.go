package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	x := 5
	val := x < 5
	fmt.Printf("%t\n", val)
	// Lägg parenteser vid flera conditions
	fmt.Println(x == 5 || x != 5 && x == 5)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What's your age?")
	scanner.Scan()
	user_age, _ := strconv.ParseInt(scanner.Text(), 10, 64)

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

	y := 5

	for y < 10 {
		y++
		fmt.Println(y)
	}

	for z := 0; z < 5; z++ {
		fmt.Println(z)
	}

	fmt.Println("Vad heter du?")
	scanner.Scan()
	name := scanner.Text()

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

// Samma jämförelser som Python. < > osv
// < > <= >= != ==
// Kan inte jämföra float och int
// AND &&
// OR ||
// NOT !
// If
// Else if
// Else
// For
// for x:=0; x<5; x++
// break & continue - samma som python
// Switch
// case x:
// default:
