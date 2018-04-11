package main

import "fmt"

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

// Don't define things like this, it's just for the presentation
// START OMIT
func (d dog) PrintFactor() {
	fmt.Printf("My pet factor is %d out of 10.\n", d.petFactor)
}

func (d dog) Bark() {
	fmt.Println("Woof woof!")
}

func (s seagull) PrintFactor() {
	fmt.Printf("My fly factor is %d out of 10.\n", s.flyFactor)
}

func main() {
	d := dog{"Snoopy", 4, 10}
	g := seagull{"Steven", 2, 8}

	d.Speak()
	g.Speak()
	d.PrintFactor()
	g.PrintFactor()
	d.Bark()
}

// END OMIT
