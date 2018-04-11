package main

import "fmt"

// START OMIT
type animal struct {
	name string
	age  int
}

func (a animal) Speak() {
	fmt.Printf("I am %s, an animal of %d years old\n", a.name, a.age)
}

type dog struct {
	animal    // This is embedding
	petFactor int
}

/*func (d dog) Speak() {
	fmt.Printf("Woof! I am %s Dog, I am %d years old and have a pet factor of %d/10\n", d.name, d.age, d.petFactor)
}*/

func main() {
	snoopy := dog{
		animal: animal{
			name: "Snoopy",
			age:  4,
		},
		petFactor: 10,
	}
	snoopy.Speak()
	snoopy.animal.Speak()
}

// END OMIT
