package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}                // Slice
	m := map[string]int{"Arjan": 123, "Ivan": 456} // Map

	fmt.Println(s[3])
	fmt.Println(m["Arjan"])
}
