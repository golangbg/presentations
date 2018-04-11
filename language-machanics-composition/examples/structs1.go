package main

import "fmt"

// START OMIT
type animal struct {
	name string
	age  int
}

// Methods are defined through a reflector
func (a animal) greet() {
	fmt.Printf("I am %s and I am %d years old", a.name, a.age)
}

func main() {
	dog := animal{name: "Snoopy", age: 4}
	dog.greet()
}

//END OMIT
