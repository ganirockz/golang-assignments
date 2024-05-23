package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel for signaling task completion
	done := make(chan bool)

	// Start the task in a goroutine
	go performTask(done)

	// Wait for either task completion or timeout
	select {
	case <-done:
		fmt.Println("Task completed successfully.")
	case <-time.After(3 * time.Second):
		fmt.Println("Task timed out.")
	}
}

func performTask(done chan<- bool) {
	// Simulate a long-running task
	time.Sleep(4 * time.Second)

	// Signal task completion
	done <- true
}
