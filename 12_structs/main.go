package main

import (
	"fmt"
)

// Unlike other languages(Ruby, JS etc..) there is no class concept in Go,
// instead we can use stucts to create instances and stuff

// Defining struct
type Person struct {
	FirstName string
	LastName  string
	Location  string
}

// add function to above Person struct => this is like public method available on Person object

func (p Person) getFullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	// instantiating struct
	person1 := Person{
		FirstName: "Yugesh",
		LastName:  "Palvai",
	}

	person1.Location = "IND"

	fmt.Println("Person object", person1)
	fmt.Println("Person1 firstName", person1.FirstName)
	fmt.Println("Person1 full name is - ", person1.getFullName())
}
