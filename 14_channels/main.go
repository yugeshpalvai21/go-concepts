package main

import (
	"fmt"
	"time"
)

// chaneels are used for communication and synchronization between goroutines
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for {
		data, ok := <-ch

		if !ok {
			break
		}

		fmt.Println("Received:", data)
	}
}

func main() {
	dataChannel := make(chan int) // create an integer channel

	go sendData(dataChannel)    // start the sendData goroutine
	go receiveData(dataChannel) // start the receiveData goroutine

	time.Sleep(time.Second * 3)
	fmt.Println("Main goroutine exiting...")
}
