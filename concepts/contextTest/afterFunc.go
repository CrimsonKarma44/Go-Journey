package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func AfterFunc() {
	// Create a context that cancels after 500ms
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// Register a function to run when the context ends
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("1. Context cancelled! Cleaning up...")
		wg.Done()
	})

	// Ideally, you defer stop() to clean up the watcher if main finishes early.
	// If stop() returns false, it means the function has already started or ran.
	defer stop()

	fmt.Println("2. Work started...")
	
	// Wait for the cleanup to happen
	wg.Wait()
	fmt.Println("3. Program finished.")
}