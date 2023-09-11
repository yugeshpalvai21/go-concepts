package main

import (
	"fmt"
	"time"
)

//GO concurrency model is based on goroutines and channels
// Goroutines are lightweight threads of execution

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 500) // Sleep for 500 milliseconds
	}
}

func main() {
	go sayHello()               // start a new goroutine
	time.Sleep(time.Second * 2) // sleep for 2 seconds in the main goroutine
	fmt.Println("Main goroutine exiting...")
}
