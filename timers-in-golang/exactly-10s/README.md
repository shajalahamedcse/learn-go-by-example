
Create A Golang Timer Using A Goroutine
Posted on February 10, 2016 by Aaron Nordyke

One of our engineering teams needed to monitor long-running jobs in some of our Golang code. If the job ran for more than 60 seconds, we would be notified. Our solution was to create a timer utility, using Go’s time package and a goroutine.
Timer Pattern

First, the basic pattern for creating a timer:

package main

import (
    "time"
    "fmt"
)

const time_in_seconds = 60

func main() {
    // 60 second timer.
    timer := time.NewTimer(time.Second*time_in_seconds)
    // Stop the timer at the end of the function. 
    // Defers are called when the parent function exits.
    defer timer.Stop()
    
    // Wait for timer to finish in an asynchronous goroutine
    go func() {
        // Block until timer finishes. When it is done, it sends a message 
        // on the channel timer.C. No other code in 
        // this goroutine is executed until that happens.
        <-timer.C
        // If main() finishes before the 60 second timer, we won't get here
        fmt.Printf("Congratulations! Your %d second timer finished.", time_in_seconds)
    }
    
    executeTheCodeThatMightTakeMoreThanSixtySeconds()
    
    // The main function is finished, so the timer will be told to stop.
}

Extract Timer Utility Function

We extracted the above pattern into a utility function. It runs a timer for a specified period of time, and then calls a supplied function.

    package main

    import (
    "time"
    "fmt"
    )

    const time_in_seconds = 60

    func main() {

    // Run timer for 60 seconds. When timer expires, 
    // call a function to print out that our timer is done.
    timer := NewTimer(time_in_seconds, func() {
        fmt.Printf("Congratulations! Your %d second timer finished.", time_in_seconds)
    })
    defer timer.Stop()
    
    
    executeTheCodeThatMightTakeMoreThanSixtySeconds()
    }

    // NewTimer creates a timer that runs for a specified number of seconds. 
    // When the timer finishes, it calls the action function.
    // Use the returned timer to stop the timer early, if needed.
    func NewTimer(seconds int, action func()) *time.Timer {
    timer := timer.NewTimer(time.Seconds * time.Duration(seconds))
    
    go func() {
        <-timer.C
        action()
    }
    
    return timer
    }

time.AfterFunc

We eventually discovered time already has a utility function, AfterFunc, that accomplishes what we did in our own utility. Fool of a Took!

timer := time.AfterFunc(time.Second*60, func() {
    fmt.Printf("Congratulations! Your %d second timer finished.", time_in_seconds)
})
defer timer.Stop()

I’ve created a working example on the Go playground: https://play.golang.org/p/wXI-U5XgnG.
Summary

Use the standard time.AfterFunc utility – not ours – But I hope these examples shed light on goroutines, channels, and defers. They are some of the coolest features in Go.
