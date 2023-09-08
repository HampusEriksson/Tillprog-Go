/*
Lektion 1

Variabler, datatyper, implicit, explicit, print, fmt, console input, type conversion, operators
*/

package main

import (
	"fmt"
)

func main() {

	// Println
	fmt.Println("Hello world")

	// Variabler
	// Static typing - variabel kan inte ändra datatyp
	// var variabel_namn datatyp = värde
	// Explicit - bestämma typen direkt
	var student string = "Lion"
	fmt.Println(student)

	// Walrus - skippa var - rekommenderas pga lättast
	// Samma sak som att skriva
	// var name string = "Hampus"
	// Implicit - låta go "välja" typen
	name := "Hampus"

	// Printf
	// %v
	// Första v = name
	// Andra v = student
	fmt.Printf("Välkommen %v %v\n", name, student)
	fmt.Println("Välkommen", name, student)
	fmt.Print("Välkommen", name, student)

	// Console input
	// Det finns flera sätt. Förra året visade jag .Scan()
	// Type conversion
	var name2 string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name2)
	fmt.Println("Hello", name2)

	// Operators
	/*
		+ - / * %
		Kan inte räkna med t.ex. float+int
		import "math"
	*/

	fmt.Println(5 + 10)
}

// Datatyper
// integer, float, string
