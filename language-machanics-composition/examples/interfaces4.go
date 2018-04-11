package main

import "fmt"

// START OMIT
type score int

func (s *score) Final() {
	fmt.Printf("The final score is %d", int(*s))
}

func main() {
	score(100).Final()
	// A pointer receiver requires and address, but in the example above there is no address
}

// END OMIT
