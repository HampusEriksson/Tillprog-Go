// https://www.youtube.com/watch?v=2ybLD6_2gKM&ab_channel=LowLevelLearning

package main

import (
	"fmt"
	"strings"
)

func changelist(myslice []int) {
	myslice[0] = 100
}

func changestring(mystring string) {
	mystring = "Hampus"
}

func check_username(username *string) {
	fmt.Println(*username)

	if strings.Contains(*username, "fan") {
		*username = "Hello_Kitty"
	}

}
func main() {
	// slice - exempel
	// slice och maps beter sig "konstigt" - de pekar på en minnesplats
	// Mutable - kan ändras
	var x []int = []int{1, 2, 3}
	y := x // y och x pekar på samma slice
	changelist(x)
	fmt.Println(y, x)

	// Immutable - kan inte ändras
	var liar string = "Johan"
	changestring(liar)
	fmt.Println(liar)

	var username string = "fansatanmolarki"
	check_username(&username)
	fmt.Println(username)

	// int - exempel
	var a int = 5
	b := a
	b = 100
	fmt.Println(a, b)
	fmt.Println(&a)

	i := 100
	fmt.Println(&i)
	z := &i
	*z = 8
	fmt.Println(i, &i, z, &z)
}
