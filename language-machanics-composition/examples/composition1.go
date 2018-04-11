package main

import "fmt"

// START OMIT
type dog struct {
	name      string
	age       int
	petFactor int
}

func (d dog) Speak() {
	fmt.Printf("Woof! I am %s Dog, I am %d years old.\n", d.name, d.age)
}

type seagull struct {
	name      string
	age       int
	flyFactor int
}

func (s seagull) Speak() {
	fmt.Printf("Mew! Mew! I am %s Seagull, I am %d years old.\n", s.name, s.age)
}

func main() {
	d := dog{"Snoopy", 4, 10}
	g := seagull{"Steven", 2, 8}

	d.Speak()
	g.Speak()
}

// END OMIT
