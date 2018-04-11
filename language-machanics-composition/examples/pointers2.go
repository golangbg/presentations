package main

import "fmt"

// START OMIT
func addTwo(i int) int {
	fmt.Printf("addTwo: i is %d, the address of i is %v\n", i, &i)
	return i + 2
}

func main() {
	i := 3
	fmt.Printf("Main: i is %d, the address of i is %v\n", i, &i)

	i = addTwo(i)
	fmt.Printf("Main: i is %d, the address of i is %v\n", i, &i)
}

// END OMIT
