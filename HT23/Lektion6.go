// Structs and Custom types

package main

import "fmt"

type Person interface {
	changeSchool(newSchool *School)
}

type Student struct {
	name   string
	age    int
	school *School
}

type Teacher struct {
	name   string
	age    int
	school *School
}
type School struct {
	name string
}

// Metoder skrivs genom att:
// func (variabelnamn Struct-typ) funktionsnamn() returtyp
func (s Student) over18() bool {
	if s.age >= 18 {
		return true
	} else {
		return false
	}
}

// En metod som ändrar på struct måste ha pointers
func (s *Student) changeSchool(newSchool *School) {
	fmt.Println(s.name, "byter skola till", newSchool.name)
	s.school = newSchool
}

func (s *Teacher) changeSchool(newSchool *School) {
	fmt.Println(s.name, "byter skola till", newSchool.name)
	s.school = newSchool
}
func main() {
	school1 := &School{"NTI Vetenskapsgymnasiet Solna"}
	school2 := &School{"KTH"}
	student1 := &Student{"Ludwig", 18, school1}
	student2 := &Student{"Leon", 17, school1}
	teacher1 := &Teacher{"Daniel", 62, school1}
	student1.name = "Hampus"
	fmt.Println(student1.school.name)
	fmt.Println(student2.over18())
	School_persons := []Person{student1, student2, teacher1}
	fmt.Println(School_persons)
	for _, person := range School_persons {
		person.changeSchool(school2)
	}

}
