## Channels in Go 

As a beginner, it is quite easy and straightforward to write code in GO. It remains that until you started writing code using `channel`.At first glance, everything about ` channels ` seems confusing and unintuitive. That means ` channels ` is one concept that we have to spend some time to master them.

Let's dive in and try to understand what is `channel` is. 

### Goroutines

If we want to understand channels properly, it is absolutly necessary to know `Goroutines` first.

Let’s start with a simple Goroutine, that takes a number, adds two with it, and at last prints its value 

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

We can think of this program that it have two blocks: one being the main funciton, and the other being the addTwo goroutine.

The problems with this implementation is that these two parts of our code are rather disconnected. As a consequence :

We cannot access the result of multiplyByTwo in the main function.
We have no way to know when the multiplyByTwo goroutine completes. As a result of this, we have to pause the main function by calling time.Sleep, which is a hacky solution at best.