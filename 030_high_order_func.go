package main

import "fmt"

type Iterable interface {
	Range(f func(int, any))
}

func Sum(s Iterable) (r any) {
	s.Range(func(i int, v any) {
		r += v
	})
	return
}

type ISlice []int

func (s ISlice) Range(f func(int, any)) {
	for i, v := range s {
		f(i, v)
	}
}

type FSlice []float32

func (s FSlice) Range(f func(int, any)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	is := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", is, Sum(is))

	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", fs, Sum(fs))
}
