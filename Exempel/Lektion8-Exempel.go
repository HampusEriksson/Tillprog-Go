package main

import (
	"fmt"
	"math/rand"
	"time"
)

func foodOrder(s string, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	students := rand.Intn(20) + 130
	fmt.Printf("%s orders food for %v students.\n", s, students)
	ch <- students
}

func main() {
	// Skapa en slice med några skolor
	schools := []string{"NTI", "LBS", "Procivitas", "Anna Whitlock", "Blackeberg", "Bromma", "Cyber", "Didaktus", "Distra", "Enskilda", "Globala", "Fryshuset", "Grillska", "Jensen", "Klara", "Kulturama", "Kungsholmen", "Kunskapsgymnasiet", "Kärrtorp"}

	// Skapar en channel för att ta emot matbeställningar
	orders := make(chan int)

	// Skapar en variabel för att summera antal ordrar
	nr_of_orders := 0

	// För varje skola startas en gorutine för funktionen foodOrder
	// _ används om inte index är relevant. som enumerate i python
	// variabeln school kommer anta varje string i slicen schools
	for _, school := range schools {
		go foodOrder(school, orders)
	}

	for _ = range schools {
		// Denna rad kommer pausa programmet tills den mottar något från kanalen
		nr_of_orders += <-orders
	}

	fmt.Printf("A total of %v portions of food is needed today.", nr_of_orders)
}
