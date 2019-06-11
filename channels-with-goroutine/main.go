package main

import "fmt"

func main() {

	num := 12

	// Here we will make a channel,
	// which will be used to move int datatype between goroutine and main function
	ch := make(chan int, 1)

	// Now we will run this function as a goroutine,
	// and will pass the channel
	go addTwo(num, ch)

	// If we receive any output on this channel, print it to the console
	fmt.Println(<-ch)
}

// This function accepts a channel as its second agrument
func addTwo(num int, ch chan<- int) {
	result := num + 2

	ch <- result
}
