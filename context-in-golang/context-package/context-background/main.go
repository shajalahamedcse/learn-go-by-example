package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	waitForMe(ctx, 5*time.Second, "I have waited 5 seconds. Now I have to go.")
}

func waitForMe(ctx context.Context, d time.Duration, s string) {

	select {

	// time.After(d) returns a boolean after some duration
	case <-time.After(d):
		fmt.Println(time.After(d))

	// ctx.Done()
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
