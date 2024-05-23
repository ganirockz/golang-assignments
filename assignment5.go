package main

import (
	"fmt"
)

func main() {

	numbers := make(chan int)
	squared := make(chan int)
	b := make(chan bool)

	// Start the pipeline stages
	go generateNumbers(numbers)
	go squareNumbers(numbers, squared)
	go printSquaredNumbers(squared, b)

	<-b

}

func generateNumbers(out chan<- int) {
	// Generate numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func squareNumbers(in <-chan int, out chan<- int) {
	// Square numbers received from the input channel
	for num := range in {
		out <- num * num
	}
	close(out)
}

func printSquaredNumbers(in <-chan int, b chan bool) {
	// Print squared numbers received from the input channel
	for squared := range in {
		fmt.Println(squared)
	}
	b <- true
}
