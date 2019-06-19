// Scenerio 1:
// Your update fuction runs each second,

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Sleep based. UGLY
	for i := 0; i < 3; i++ {
		time.Sleep(5 * time.Second)
		update("Sleep")
	}
	// Sleep based. VERY BAD
	for i := 0; i < 3; i++ {
		<-time.After(5 * time.Second)
		update("Channel")
	}
	// Shot and forget. Will work but bad practice. it may crash, if you are not sure about the execution time of your function: BAD
	for i := 0; i < 3; i++ {
		<-time.After(5 * time.Second) // Execution time must be lower than this value
		go update("GoRoutine")
	}
	// Take care of yourself. Don't shot and forget. This is the Good practice so far. Perfect example of defensive coding :p GOOD
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		<-time.After(5 * time.Second) // Execution time must be lower than this value
		wg.Add(1)
		go func() {
			defer wg.Done()
			update("GoRoutine")
		}()
		wg.Wait()
	}
	// A bit more defense. Here I'm using my lib. You can implement your own defense mechanism. BEST
	q := New(1)
	for i := 0; i < 3; i++ {
		<-time.After(5 * time.Second) // Execution time must be lower than this value
		q.Add()
		go func() {
			defer q.Done()
			update("Work queue")
		}()
		q.Wait()
	}

}

// A complex function that take random time to run
// it may take 5s to 10s or whatever
func update(s string) {
	fmt.Println("Hello from", s)
	// Unfortunately A function took longer
	// Don't worry, Just for simulation
	time.Sleep(3 * time.Second)
	fmt.Println("Finish from", s)
}
