package main

import (
	"fmt"
	"math"
)

// Skapa en funktion
// Parametrar skrivs som (variabelnamn datatyp)
// (variabelnamn datatyp, variabelnamn datatyp)
// (variabelnamn, variabelnamn datatyp)
func pythagoras(a, b float64) float64 {
	//math.Pow(2, 10)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// behöver inte säga return x,y
func sums(a, b int) (sum int) {
	// defer - görs i slutet
	defer fmt.Println("Nu är funktionen returnerad")
	fmt.Println("Vi ska summera", a, b)
	sum = a + b
	return
}

// closures

func check_runs() func() {
	runs := 0

	return func() {
		runs++ // ++ ökar variabeln med 1
		fmt.Println(runs)
	}
}

// Funktionen main körs när filen körs
func main() {
	fmt.Println("Detta körs när filen körs.")
	// Kalla på en funktion

	fmt.Println(pythagoras(2, 3))
	fmt.Println(sums(5, 6))

	// funktion är en datatyp. kan t.ex. sparas i variabel eller returneras
	test := func() {
		fmt.Println("Testing")
	}

	test()
	test()

	get_my_runs := check_runs()
	get_my_runs()
	get_my_runs()
	get_my_runs()

}
