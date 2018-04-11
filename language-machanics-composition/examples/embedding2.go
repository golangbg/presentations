package main

import "fmt"

// START OMIT
type animal struct {
	name string
	age  int
}

type dog struct {
	animal    // This is embedding
	petFactor int
}

func main() {
	snoopy := dog{
		animal: animal{
			name: "Snoopy",
			age:  4,
		},
		petFactor: 10,
	}
	fmt.Println(snoopy.name) // Access through the outer struct
}

// END OMIT
