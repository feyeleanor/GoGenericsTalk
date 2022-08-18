package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case []T:
		for _, x := range c {
			r = f(r, x)
		}
	case map[int]T:
		for _, x := range c {
			r = f(r, x)
		}
	case chan T:
		for x := range c {
			r = f(r, x)
		}
	}
	return
}

func Pump[T Numeric](ic chan<- T, v ...T) {
	for _, v := range v {
		ic <- v
	}
	close(ic)
}
func main() {
	is := []int{0, 1, 2, 3, 4}
	ir := Reduce(is, func(x, v int) int {
		return x + v
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", is, ir, ir)

	im := map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}
	ir = Reduce(im, func(x, v int) int {
		return x + v
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", im, ir, ir)

	ic := make(chan int)
	go Pump(ic, 0, 1, 2, 3, 4)
	ir = Reduce(ic, func(x, v int) int {
		return x + v
	})
	fmt.Printf("(chan ic).Reduce(f()) = %v [%T]\n", ir, ir)
}
