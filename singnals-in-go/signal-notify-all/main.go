package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	ch := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(ch)

	// Block until any signal is received.
	s := <-ch
	fmt.Println(s)
}
