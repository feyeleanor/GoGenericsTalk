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

type ISlice []int

func (s ISlice) Range(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v\n", is, Sum(is))
}
