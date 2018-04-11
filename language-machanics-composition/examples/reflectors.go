package main

import (
	"fmt"
	"strconv"
)

type strint int

func (s strint) String() string {
	return strconv.Itoa(int(s))
}

func main() {
	a := strint(7)
	fmt.Printf("a as string: %s", a)
}
