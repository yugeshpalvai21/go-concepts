package main

import (
	"fmt"
)

func main() {
	// pointers in Go are just variables that stores memory address of another variable

	var ptr *int
	x := 10

	ptr = &x

	fmt.Println("memorey address that pointer holds", ptr)
	fmt.Println("printing values of x and pointer", x, *ptr)

	// if we change the value of pointer that changes the value of x too..

	*ptr = 20
	fmt.Println("printing values of x and pointer", x, *ptr)
}
