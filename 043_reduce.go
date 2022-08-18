package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Reducible[T Numeric] interface {
	Reduce(func(T, T) T) T
}

type ISlice []int

func (s ISlice) Reduce(f func(x, v int) int) (r int) {
	for _, v := range s {
		r = f(r, v)
	}
	return
}

type FArray [3]float32

func (a FArray) Reduce(f func(x, v float32) float32) (r float32) {
	for _, v := range a {
		r = f(r, v)
	}
	return
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case Reducible[T]:
		r = c.Reduce(f)
	}
	return
}

func DoReduce[T Numeric](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

func main() {
	DoReduce(ISlice{0, 1, 2, 3, 4}, func(x, v int) int {
		return x + v
	})

	DoReduce(FArray{0, 1, 2}, func(x, v float32) float32 {
		return x + v
	})
}
