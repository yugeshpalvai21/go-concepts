package main

import (
	"fmt"
)

// regular fizzbuzz programm where print
// 'fizz' => number divisible by 3
// 'buzz' => number divisible by 5
// 'fizzbuzz' => number divisible by both 3 and 5
// number => if not satisfy above conditions

func fizzbuzzWithLoops(num1, num2 int) {
	for i := num1; i <= num2; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzbuzzWithSwitch(num1, num2 int) {
	for i := num1; i <= num2; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzzWithLoops(100, 150)
	fizzbuzzWithSwitch(10, 50)
}
