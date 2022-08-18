package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Reducible[T Numeric] []T

func (s Reducible[T]) Reduce(f func(T, T) T) (r T) {
	for _, x := range s {
		r = f(r, x)
	}
	return
}

type ISlice = Reducible[int]

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	ir := is.Reduce(func(x, v int) int {
		return x + v
	})
	fmt.Printf("(%T)%v.Reduce(f()) = %v [%T]\n", is, is, ir, ir)

	fs := Reducible[float32]{0, 1, 2, 3, 4}
	fr := fs.Reduce(func(x, v float32) float32 {
		return x + v
	})
	fmt.Printf("(%T)%v.Reduce(f()) = %v [%T]\n", fs, fs, fr, fr)
}
