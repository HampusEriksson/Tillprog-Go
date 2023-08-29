package main

import "fmt"

//"bufio"
//"os"
//"strconv"

func main() {
	// Arrays/slices  - ordered collection of items
	// Element  - saker i en array
	// lista i python

	// Tom slice
	var numbers []int
	fmt.Println(numbers)

	// Slice med start-element
	// variabelnamn := []int{elements}
	students := []string{"Adam", "Beatrice", "Cedric", "Daniella"}

	// append, lite annorlunda än Python
	// Vi slår ihop en lista med ett element och sparar i samma variabel
	students = append(students, "Erik")
	// students.append("Erik") i Python
	fmt.Println(students)

	// len = längden av en slice
	fmt.Println(len(students))

	// loopa igenom slice
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}
	// for index, element := range arraynamn
	for index, student := range students {
		if student != "" {
			fmt.Println(index, ":", student)
		}

	}
	var indexToRemove int
	fmt.Print("Enter the index of the student to remove: ")
	fmt.Scan(&indexToRemove)

	// Remove the element at the specified index
	students = append(students[:indexToRemove], students[indexToRemove+1:]...)

	// Slices har inte fast size - Arrays har fast size
	// It's a good idea to use slices unless you have a specific reason to use arrays, such as one of the reasons mentioned above.
	// var variabelnamn [maxsize]datatyp
	var deck [60]string

	fmt.Println(deck)

	// capacity
	fmt.Println(cap(deck))

	// variabelnamn := make([]datatyp, maxsize)
	grades := make([]int, 5)
	fmt.Println(grades)

	// Maps i GO = Dictionaries i Python
	// var accounts map[string]string = map[string]string{}
	var accounts map[string]string = map[string]string{
		"Amanda":  "Amanda123",
		"Bo":      "Bo123",
		"Cecilia": "Cecilia123",
	}

	fmt.Println(accounts)
	fmt.Println(accounts["Amanda"])

	// delete(map, key)
	delete(accounts, "Cecilia")

	// value, ok := mp[key]
	value, ok := accounts["Amanda"]
	fmt.Println(value, ok)

	if !ok {
		fmt.Println("Finns ej")
	}

}
