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

	// Printf
	// %v
	// Variabler
	// Static typing - variabel kan inte ändra datatyp
	// var variabel_namn datatyp = värde

	// Behöver inte ge värde direkt

	// Explicit - bestämma typen direkt
	// Implicit - låta go "välja" typen

	// Walrus - skippa var - rekommenderas pga lättast

	// Console input
	// Det finns flera sätt. Förra året visade jag .Scan()
	// Type conversion
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)

	// Operators
	/*
		+ - / * %
		Kan inte räkna med t.ex. float+int
		import "math"
	*/
}

// Datatyper
// integer, float, string
