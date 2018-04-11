package main

import (
	"fmt"
	"strings"
)

type animal struct {
	name string
	age  int
}

// START OMIT

// Value semantics
func (a animal) lowercaseName() {
	a.name = strings.ToLower(a.name)
}

// Pointer semantics
func (a *animal) uppercaseName() {
	a.name = strings.ToUpper(a.name)
}

func main() {
	dog := animal{name: "Snoopy", age: 4}

	dog.lowercaseName()
	fmt.Printf("After lowercaseName, dog.name is: %s\n", dog.name)

	dog.uppercaseName()
	fmt.Printf("After uppercaseName, dog.name is: %s\n", dog.name)
}

//END OMIT
