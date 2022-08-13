package main

import "fmt"

type Iterable interface {
	Range(f func(int, int))
}

func Sum(s Iterable) (r int) {
	s.Range(func(i, v int) {
		r += v
	})
	return
}

type FSlice []float32

func (s FSlice) Range(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", fs, Sum(fs))
}
