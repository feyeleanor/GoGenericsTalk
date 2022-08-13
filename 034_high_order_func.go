package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Indexable[T any] interface {
	~[]T
}

func Reduce[T Indexable[E], E Numeric](s T, f func(E, E) E) (r E) {
	for _, v := range s {
		r = f(r, v)
	}
	return
}

func main() {
	is := []int{0, 1, 2, 3, 4}
	ir := Reduce(is, func(x, v int) int {
		return x + v
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", is, ir, ir)

	fs := []float32{0, 1, 2, 3, 4}
	fr := Reduce(fs, func(x, v float32) float32 {
		return x + v
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", fs, fr, fr)
}
