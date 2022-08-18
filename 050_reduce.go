package main

import "fmt"

func Reduce[T any](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case interface{ Range(f func(int, T)) }:
		c.Range(func(i int, v T) {
			r = f(r, v)
		})
	case interface{ Reduce(func(T, T) T) T }:
		r = c.Reduce(f)
	}
	return
}

func DoReduce[T any](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

type ISlice []int

func (s ISlice) Range(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

type FArray [3]float32

func (a FArray) Reduce(f func(x, v float32) float32) (r float32) {
	for _, v := range a {
		r = f(r, v)
	}
	return
}

func main() {
	DoReduce(ISlice{0, 1, 2, 3, 4}, func(x, v int) int {
		return x + v
	})

	DoReduce(FArray{0, 1, 2}, func(x, v float32) float32 {
		return x + v
	})
}
