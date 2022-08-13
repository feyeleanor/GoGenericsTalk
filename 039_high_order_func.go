package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Reducible[T Numeric] chan T

func (c Reducible[T]) Reduce(f func(T, T) T) (r T) {
	for x := range c {
		r = f(r, x)
	}
	return
}

func (c Reducible[T]) Pump(v ...T) {
	for _, v := range v {
		c <- v
	}
	close(c)
}

type IChan = Reducible[int]
type FChan = Reducible[float32]

func main() {
	ic := make(IChan)
	go ic.Pump(0, 1, 2, 3, 4)
	ir := ic.Reduce(func(x, v int) int {
		return x + v
	})
	fmt.Printf("(chan ic).Reduce(f()) = %v [%T]\n", ir, ir)

	fc := make(FChan)
	go fc.Pump(0, 1, 2, 3, 4)
	fr := fc.Reduce(func(x, v float32) float32 {
		return x + v
	})
	fmt.Printf("(chan fc).Reduce(f()) = %v [%T]\n", fr, fr)
}
