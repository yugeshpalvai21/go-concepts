package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}

	return a / b, nil
}

func main() {
	result, err := divide(1, 0) // change these arguments to see different logs on console
	if err != nil {
		fmt.Println("Error Occured: ", err)
	} else {
		fmt.Println("result of divide function call", result)
	}
}
