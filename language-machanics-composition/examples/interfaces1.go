package main

import "fmt"

// START OMIT
type speaker interface {
	Speak()
}

type factorPrinter interface {
	PrintFactor()
}

type barker interface {
	Bark()
}

// Function using an interface as argument
func makeBark(b barker) {
	b.Bark()
}

func main() {
	d := dog{"Snoopy", 4, 10}
	makeBark(d)
}

// END OMIT

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

func (d dog) PrintFactor() {
	fmt.Printf("My pet factor is %d out of 10.\n", d.petFactor)
}

func (d dog) Bark() {
	fmt.Println("Woof woof!")
}

func (s seagull) PrintFactor() {
	fmt.Printf("My fly factor is %d out of 10.\n", s.flyFactor)
}
