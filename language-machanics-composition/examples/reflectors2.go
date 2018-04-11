package main

import "fmt"

type animal struct {
	name string
	age  int
}

// START OMIT
// Value receiver
func (a animal) printName() {
	fmt.Println(a.name)
}

// Pointer receiver
func (a *animal) showName() {
	fmt.Println(a.name)
}

func main() {
	a := animal{"Snoopy", 4}

	// In reality just syntactical sugar
	a.printName()
	a.showName()

	// What actually happens under the hood
	(animal).printName(a)
	(*animal).showName(&a)
}

// END OMIT
