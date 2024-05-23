package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the value of N : ")
	fmt.Scanln(&n)

	c := incrementor(n)
	sum := sumIt(c)

	fmt.Println(sum)

}

func incrementor(n int) chan int {
	out := make(chan int)

	go func() {
		for i := 0; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func sumIt(c chan int) int {
	sum := 0
	func() {
		for n := range c {
			//fmt.Println("received ", n)
			sum += n
		}
	}()

	return sum
}
