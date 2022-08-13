package main

import "fmt"

func main() {
	var c chan int

	c = make(chan int, 16)
	go func(c chan<- int, v ...int) {
		for _, n := range v {
			fmt.Println("sending:", n)
			c <- n
		}
		close(c)
	}(c, 0, 2, 4, 6)

	for n := range c {
		fmt.Println("n:", n)
	}

	fmt.Println("n:", <-c)
}
