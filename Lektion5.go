// https://www.youtube.com/watch?v=2ybLD6_2gKM&ab_channel=LowLevelLearning

package main

import "fmt"

func main() {
	toChange := "Hello"
	var pointer *string = &toChange
	fmt.Println(pointer, &pointer, *pointer)
	*pointer = "Good bye"
	fmt.Println(pointer, &pointer, *pointer)

}

// Mutable - kan ändras
// int - exempel
// slice - exempel
/*var x []int = []int{3, 4, 5}
y := x // y och x pekar på samma slice
y[0] = 100
fmt.Println(y)*/
// slice och maps beter sig "konstigt" - de pekar på en minnesplats
// exempel - skicka in slice i en funktion
// array skapar en kopia
// Immutable
// Pointers
// fmt.Println(&x)
// y:= &x
// *y = 8
