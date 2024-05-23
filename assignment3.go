package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func main() {
	numGoroutines := 5

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go incrementCounter(i)
	}

	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

func incrementCounter(id int) {
	defer wg.Done()

	mutex.Lock()
	counter++
	fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)
	mutex.Unlock()
}
