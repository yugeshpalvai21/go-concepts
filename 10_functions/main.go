package main

import (
	"fmt"
)

// function with two parameters and return type
func addNums(num1, num2 int) int {
	return num1 + num2
}

// function with dynmic parameters and return type
func getDynamicSum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//named return values

func divide(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = fmt.Errorf("Division by Zero")
		return // return result as nil and err as error response
	}
	result = num1 / num2
	return // return result and err as nil
}

func main() {
	// sum of two numbers
	num1, num2 := 10, 11
	sum := addNums(num1, num2)
	fmt.Println(sum)

	//dynamic sum of numbers
	dynamic_sum := getDynamicSum(1, 2, 3, 5, 6) //these values converted into array of elements
	fmt.Println(dynamic_sum)

	// calling named return value function
	result, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// second time calling above function with valid args
	quotiant, err := divide(12.0, 13.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotiant)
	}
}
