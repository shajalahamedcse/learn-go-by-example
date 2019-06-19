package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Timeout example

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	waitForMe(ctx, 5*time.Second, "I waited 5 second for you . Now I have to go.")
}

func waitForMe(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
