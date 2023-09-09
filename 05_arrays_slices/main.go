package main

import (
	"fmt"
	"strings"
)

func main() {
	// declaration of array => fixed lenght of elements
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println("Array One is", arr1)

	//shorthand syntax
	arr2 := [3]string{"first", "second", "third"}
	fmt.Println("Array Two is", arr2)

	//declaring slice => dynamic length of elements
	arr3 := []int{100, 101, 102, 105}
	fmt.Println("SLice values", arr3)

	// print types of each variable
	fmt.Printf("%T %T %T \n", arr1, arr2, arr3)

	fmt.Println("Lenght of arr3 is", len(arr3))

	//append elements to slice
	arr3 = append(arr3, 1001)
	fmt.Println("New Array", arr3)

	//delete element from slice
	index := 3
	arr3 = append(arr3[:index], arr3[index+1:]...)
	fmt.Println("Array with deleted element", arr3)

	//join slice elements into string
	fmt.Println("Joined string", strings.Join(arr2[:], "+"))

	//capacity of arr
	fmt.Println("Capacity of arr1", cap(arr1))
}
