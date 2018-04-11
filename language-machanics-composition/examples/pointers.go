package main

import "fmt"

// START OMIT
// Value semantics
func addTwo(i int) int {
	return i + 2
}

// Pointer semantics
func multiplyTwo(i *int) {
	*i *= 2
}

func main() {
	i := 3
	fmt.Printf("i is %d\n", i)

	i = addTwo(i) // We're passing the VALUE of i
	fmt.Printf("i is %d\n", i)

	multiplyTwo(&i) // We're passing the ADDRESS of i
	fmt.Printf("i is %d\n", i)
}

// END OMIT
