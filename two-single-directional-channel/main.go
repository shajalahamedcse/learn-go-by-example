package main

import "fmt"

func main() {
	num := 12

	in := make(chan int)
	out := make(chan int)

	// We supply 2 channels to the `addTwo` function
	// One for sending data and another for receiving
	go addTwo(in, out)

	// Now we send the data into `in` channel
	// and wait for the result from `out` channel
	in <- num

	fmt.Println(<-out)
}

func addTwo(in <-chan int, out chan<- int) {

	fmt.Println("Initializing goroutine...")

	// The goroutine does not proceed until data is received on the `in` channel
	num := <-in

	result := num + 2
	out <- result
}
