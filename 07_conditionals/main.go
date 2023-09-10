package main

import (
	"fmt"
)

func main() {
	num1, num2 := 101, 102

	// if else conditionals
	if num1 > num2 {
		fmt.Printf("%d is greater than %d", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d is less than %d", num1, num2)
	} else {
		fmt.Printf("%d is equal to %d", num1, num2)
	}

	// there is no ternary operator in go instead we can use if else condition

	condition := true
	result := ""

	if condition {
		result = "It's true"
	} else {
		result = "It's false"
	}

	fmt.Println(result)

	// switch condition
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	case "Friday":
		fmt.Println("It's Friday")
	default:
		fmt.Println("It's other day")
	}
}
