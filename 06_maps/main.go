package main

import "fmt"

func main() {
	// maps are collection of elements and each element identified by unique key
	// creating map and assign values
	names := make(map[string]int)

	names["emp1"] = 101
	names["emp2"] = 102

	fmt.Println("Defined map >>", names)

	// shorthand syntax for map
	employee := map[string]string{"firstName": "Yugesh", "lastName": "Palvai"}

	//adding element to existing map
	employee["timeZone"] = "IST"
	fmt.Println("Employee Name", employee)

	//deleting element from existing map
	delete(employee, "timeZone")
	fmt.Println("Updated Employee", employee)

	// iterate through map
	for key, value := range employee {
		fmt.Printf("Key - %s, Value - %s ", key, value)
	}

	//check if key is available or not
	value, exists := employee["timeZone"]

	if exists {
		fmt.Println("key available", value)
	} else {
		fmt.Println("Key is not avaialble")
	}
}
