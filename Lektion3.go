package main

import "fmt"

//"bufio"
//"os"
//"strconv"

func main() {
	var arr [5]int
	fmt.Println(arr)

}

// Arrays  - ordered collection of items
// Element  - saker i en array
// var variabelnamn [maxsize]datatyp
// Index
// variabelnamn := [maxsize]int{elements}
// variabelnamn := make([]datatyp, maxsize)
// len
// loopa igenom array
// Problem med arrays - sizen bestäms direkt
// Slices löser detta
// Slices har inte fast size
// var variabelnamn []datatyp = array[start:stop]
// capacity
// appenda till en slice, spara i variabeln
// for index, element := range arraynamn
// var accounts map[string]string = map[string]string{}
// delete(map, key)
// value, ok := mp[key]
