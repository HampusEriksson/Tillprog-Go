/*

Python exempel klasser:

class Student:
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade

    def display_student_info(self):
        print(f"Name: {self.name}, Age: {self.age}, Grade: {self.grade}")

# Creating instances of the Student class
student1 = Student("Alice", 17, "A")
student2 = Student("Bob", 18, "B")

# Displaying student information
student1.display_student_info()
student2.display_student_info()

*/

// Structs and Custom types

package main

import "fmt"

// In Go, a type implicitly satisfies
// an interface if it implements all the methods
// declared by that interface.
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
func (stu Student) over18() bool {
	if stu.age >= 18 {
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
	fmt.Println(student1.school.name)
	student1.changeSchool(school2)
	fmt.Println(student1.school.name)
	fmt.Println(student1.over18())

	School_persons := []Person{student1, student2, teacher1}
	fmt.Println(School_persons)
	for _, person := range School_persons {
		person.changeSchool(school2)
	}

}
