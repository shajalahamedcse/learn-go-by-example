package main

import (
	"fmt"
	"time"
)

func main() {

	num := 12

	// We want to run a go routine to add 2 with num
	go addTwo(num)

	// Now we have to pause the program for a second so that the `addTwo` goroutine
	// can finish and print the output to console before the program exits
	time.Sleep(time.Second)
}

func addTwo(num int) int {

	result := num + 2
	fmt.Println(result)
	return result
}
