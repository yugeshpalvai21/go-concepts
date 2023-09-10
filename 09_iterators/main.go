package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	nameDetails := map[string]string{"first": "Yugesh", "last": "Palvai"}

	// we use combination of for and range to iterate over each elements in arrays/slices or maps
	for i, value := range numbers {
		fmt.Printf("value at index-%d is %d", i, value)
	}

	// if we don't want use index then we can re-write above condition like below
	for _, value := range numbers {
		fmt.Printf("value - %d", value)
	}

	// iterate elements in given map
	for key, value := range nameDetails {
		fmt.Printf("key-%s,value-%s", key, value)
	}

}
