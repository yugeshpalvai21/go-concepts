package main

import (
	"fmt"
)

func main() {
	// go support 'for' loop to run statements multiple times
	// initialization, condition, and updation in single line
	for i := 0; i < 10; i += 1 {
		fmt.Println("Number", i)
	}

	// go doesn't support while loops like other languages
	// we can acheieve same in for loop ways like initialization and conditions and updation in different areas
	n := 1
	for n < 5 {
		fmt.Println("Number", n)
		n++
	}
}
