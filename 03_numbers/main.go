package main

import "fmt"

func main() {
	// below are different types of numbers

	// integers (with negative and positives) - int, int8, int16, int32, int64
	var num1 int= 111
	var num2 int = -121

	num3 := 1100  // shorthand syntax

	fmt.Printf("%T %T %T\n", num1, num2, num3)
	
	// rune - alias for int32
	var num4 rune = 112211
	fmt.Printf("%T\n", num4)

	// integers (with only positives) - uint, uint8, uint16, uint32, uint64
	var num5 uint = 12121121
	// num5 = -12121 - we can't re-assign with negative numbers
	fmt.Printf("%T\n", num5)

	// byte - alias for uint8

	var num6 byte = 112233
	fmt.Printf("%T\n", num6)

	// floats (numbers with decimals) - float32, float64

	var num7 float32
	num7 = 121.121
	fmt.Printf("%T\n", num7)

	num8 := float64(1122.1212121)
	fmt.Printf("%T\n", num8)

	
	// really large numbers - complex64, complex128
	num9 := complex64(1 + 2i)
	num10 := complex128(3 + 4i)
	fmt.Printf("%T %T\n", num9, num10)

	// Constants - values can't be changed

	const pi float32 = 3.1415
	// pi = 12.1212 => will throw an error
	fmt.Printf("%T \n", num11)
}