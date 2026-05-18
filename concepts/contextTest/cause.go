package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Define a custom error so we can check for it later
var ErrTooSlow = errors.New("operation was too slow")
var ErrNetworkFail = errors.New("network connection lost")

func cause() {
	// 1. Create a context that supports a "Cause"
	// Notice the cancel function signature changes to accept an error.
	ctx, cancel := context.WithCancelCause(context.Background())

	// Simulate a worker
	go func() {
		// Simulate doing some work
		time.Sleep(500 * time.Millisecond)
		
		// 2. Cancel with a SPECIFIC reason
		fmt.Println("Worker: Encountered a network issue, cancelling...")
		cancel(ErrNetworkFail) 
	}()

	// Block until the context is done
	<-ctx.Done()

	// 3. Inspect the cause
	// Standard ctx.Err() is generic
	fmt.Printf("Standard Err(): %v\n", ctx.Err()) 
	
	// context.Cause() reveals the specific error you passed
	fmt.Printf("context.Cause(): %v\n", context.Cause(ctx))

	// You can now handle logic based on the specific error
	if errors.Is(context.Cause(ctx), ErrNetworkFail) {
		fmt.Println(">> Handling network failure cleanup...")
	}
}