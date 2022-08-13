package main

import "fmt"

func main() {
	c := make(chan int, 16)
	done := make(chan bool)

	go func(v ...int) {
		for _, n := range v {
			fmt.Println("sending:", n)
			c <- n
		}
		close(c)
		done <- true
	}(0, 2, 4, 6)

Receiver:
	for {
		select {
		case n, ok := <-c:
			if ok {
				fmt.Printf("len(c): %v, n: %v\n", len(c), n)
			}
		case <-done:
			break Receiver
		}
	}
}
