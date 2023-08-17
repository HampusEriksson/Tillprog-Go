package main

import "fmt"

//"bufio"
//"os"
//"strconv"

func main() {
	// Arrays  - ordered collection of items
	// Element  - saker i en arrays
	// var variabelnamn [maxsize]datatyp
	var arr [5]int
	fmt.Println(arr)
	// variabelnamn := [maxsize]int{elements}
	students := [32]string{"Jorge", "Leon", "Vigor", "Jubin"}
	fmt.Println(students)

	// variabelnamn := make([]datatyp, maxsize)
	grades := make([]int, 5)
	fmt.Println(grades)

	// len
	fmt.Println(len(grades))

	// loopa igenom array
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	// Slices har inte fast size
	// var variabelnamn []datatyp = array[start:stop]
	var favorite_students []string = students[2:4]
	fmt.Println(favorite_students)

	// capacity
	fmt.Println(cap(students))
	fmt.Println(cap(favorite_students))

	// appenda till en slice, spara i variabeln
	favorite_students = append(favorite_students, "Sam")
	fmt.Println(favorite_students)
	fmt.Println(cap(favorite_students))
	fmt.Println(students)

	// for index, element := range arraynamn
	for _, student := range students {
		if student != "" {
			fmt.Println(student)
		}

	}

	// var accounts map[string]string = map[string]string{}
	var accounts map[string]string = map[string]string{
		"Leon04":  "Leon123",
		"Jorge04": "Jorge123",
		"Mary04":  "Mary123",
	}

	fmt.Println(accounts)
	fmt.Println(accounts["Leon04"])
	// delete(map, key)
	delete(accounts, "Mary04")

	// value, ok := mp[key]
	value, ok := accounts["Leon04"]
	fmt.Println(value, ok)

	if !ok {
		fmt.Println("Finns ej")
	}

}
