package main

import (
	"fmt"
	"strings"
)

func main() {

	x := 5
	val := x < 5
	fmt.Printf("%t\n", val)
	fmt.Println(val)

	// AND && - shift + 6
	// OR || - alt gr + <
	// NOT !
	// && har prioritet över ||
	fmt.Println(x == 5 || x != 5 && x == 5)

	var name2 string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name2)

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

	// Vi kan använda for som en while-loop
	for y < 10 {
		// ökar y med ett. Samma som y +=1 i Python
		y++
		fmt.Println(y)
	}

	for z := 0; z < 15; z += 2 {
		fmt.Println(z)
	}

	//for i in range(0,15,2):
	//	print(i)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	// .Title är som capitalize i Python!
	name_cap := strings.Title(name)
	fmt.Println(name_cap)

	// Switch
	// case x:
	// default:

	switch name {

	case "Leroy":
		fmt.Println("Du gillar att gå till skolan.")

	case "Isak":
		fmt.Println("Du gillar att spela roblox.")

	case "M.a.n":
		fmt.Println("Du gillar att gå till skolan med Leroy.")

	default:
		fmt.Println("Du är inte en av de tre coola.")
	}

}
