package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello world")
	// Variabler
	// Static typing - variabel kan inte ändra datatyp
	// var variabel_namn datatyp = värde
	var fruit string = "Eggplant"
	fmt.Println(fruit)
	fruit = "Apple"

	// Behöver inte ge värde direkt
	var car int
	fmt.Println(car)

	// Explicit - bestämma typen direkt
	// Implicit - låta go "välja" typen
	var age = 28
	fmt.Println(age)
	fmt.Printf("Type: %T Value: %v", age, age)

	// Walrus - skippa var - rekommenderas pga lättast
	fabian_age := 18
	fmt.Println(fabian_age)

	// Console input
	// Type conversion
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What's your name?")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("What's your age?")
	scanner.Scan()
	user_age, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println(name, user_age)
	// Operators
	/*
		+ - / * %
		Kan inte räkna med t.ex. float+int
		import "math"
	*/
}

// Datatyper
// integer, float, string

// Printf
// Println
