package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(func(int, T))
}

func Sum[T Numeric](s Iterable[T]) (r T) {
	s.Range(func(i int, v T) {
		r += v
	})
	return
}

type ISlice []int

func (s ISlice) Range(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

type FSlice []float32

func (s FSlice) Range(f func(int, float32)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", is, Sum[int](is))

	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", fs, Sum[float32](fs))
}
