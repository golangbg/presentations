package main

import "fmt"

// START OMIT
type barker interface {
	Bark()
}

func makeBark(b barker) {
	b.Bark()
}

type dog struct {
	name string
	age  int
}

// This time we're using a pointer receiver
func (d *dog) Bark() {
	fmt.Println("Woof! Woof!")
}

func main() {
	d := dog{"Snoopy", 4}
	makeBark(d) // Will this work?
}

// END OMIT
