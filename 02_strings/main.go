package main

import (
	"fmt"
	"strings"
)

func main() {
	// traditional way of declaration
	var first_name string = "Yugesh"
	fmt.Println("First Name:", first_name)

	// shorthand syntax
	last_name := "Palvai"
	fmt.Println("last name:", last_name)

	// length of the string
	fmt.Println("Number of Chars", len(first_name))

	// convert string to uppercase, lowercase and capitalize
	fmt.Println(strings.ToUpper(first_name))
	fmt.Println(strings.ToLower(last_name))
	fmt.Println(strings.Title(first_name))

	// string concatenation

	full_name := first_name + " " + last_name
	fmt.Println(full_name)
}
