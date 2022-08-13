package main

import "fmt"

func main() {
	var c chan int

	c = make(chan int)
	go send(c, 0, 2, 4, 6)
	receive(c)
}

func send(c <-chan int, v ...int) {
	for _, n := range v {
		c <- n
	}
	close(c)
}

func receive(c chan<- int) {
	for n := range c {
		fmt.Println("n:", n)
	}
}
