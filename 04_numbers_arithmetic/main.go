package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1, num2, num3, num4 := 10, 20, 11.11, 22.22

	// arithmentic operations
	int_sum := num1 + num2
	float_sub := num4 - num3
	float_multification := num3 * num4
	int_div := num1 / num2

	fmt.Println("Addition, Sub, Div and multification", int_sum, float_sub, int_div, float_multification)

	// get random number between 1 to 100

	randomNum := rand.Intn(100)

	fmt.Println("Random Number is", randomNum)

	// get the latest time

	fmt.Println("Currrent TIme Is", time.Now())
}
