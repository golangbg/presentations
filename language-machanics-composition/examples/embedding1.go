package main

import "fmt"

// START OMIT
type animal struct {
	name string
	age  int
}

type dog struct {
	beast     animal // This is NOT embedding
	petFactor int
}

func main() {
	snoopy := dog{
		beast: animal{
			name: "Snoopy",
			age:  4,
		},
		petFactor: 10,
	}
	fmt.Println(snoopy.beast.name) // Access through the inner struct
}

// END OMIT
