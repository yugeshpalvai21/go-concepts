package main

import (
	"fmt"
)

// Defining struct
type Person struct {
	Name string
	Age  int
}

// create pointer constrictor
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// add public method to struct
func (p Person) getDetails() string {
	result := fmt.Sprintf("Name: %s, and Age: %d", p.Name, p.Age)
	return result
}

func main() {
	//instantiating struct
	person1 := Person{}

	fmt.Println("person object", person1)
	fmt.Println("person name >", person1.Name)
	// change person1 age
	person1.Age = 29
	fmt.Println("person age >", person1.Age)
	fmt.Println("person memory address", &person1)

	person2 := NewPerson("Raam", 30)
	fmt.Println("person2 object...", person2)
	fmt.Println("person2 details...", person2.getDetails())
}
