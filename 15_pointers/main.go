package main

import (
	"fmt"
)

func main() {
	var ptr *int // defining variable that stores memory address of other integer variable

	x := 5
	ptr = &x

	fmt.Println("Value of x variable -", x)
	fmt.Println("Value of ptr variable -", ptr)
	fmt.Println("Value of memory address which stored in ptr variable", *ptr)

	// if i change the value of memory address that stored in ptr variable

	*ptr = 15
	fmt.Println("Updated Value of x variable -", x)
	fmt.Println("Updated Value of memory address which stored in ptr variable", *ptr)
}
