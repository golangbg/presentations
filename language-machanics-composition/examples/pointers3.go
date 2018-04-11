package main

import "fmt"

// START OMIT
func multiplyTwo(i *int) {
	fmt.Printf("multiplyTwo: i is %d, the address of i is %v\n", *i, i)
	*i *= 2
}

func main() {
	i := 3
	fmt.Printf("Main: i is %d, the address of i is %v\n", i, &i)

	multiplyTwo(&i)
	fmt.Printf("Main: i is %d, the address of i is %v\n", i, &i)
}

// END OMIT
