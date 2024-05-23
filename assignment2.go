package main

import "fmt"

func main() {

	c := incrementor()

	//Consumer
	for n := range c {
		fmt.Println(n)
	}
}

// producer
func incrementor() chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
