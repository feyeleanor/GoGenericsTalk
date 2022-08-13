package main

import "fmt"

func main() {
	var c chan int

	c = make(chan int)
	fmt.Println("cap(c):", cap(c))
	go send(c, 0, 2, 4, 6)
	receive(c)
}

func send(c chan<- int, v ...int) {
	for _, n := range v {
		fmt.Println("sending:", n)
		c <- n
	}
	close(c)
}

func receive(c <-chan int) {
	for n := range c {
		fmt.Printf("len(c): %v, n: %v\n", len(c), n)
	}
}
