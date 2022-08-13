package main

import "fmt"

type Summable interface {
	Sum() any
}

type ISlice []int

func (s ISlice) Sum() any {
	r := 0
	for x := len(s) - 1; x > -1; x-- {
		r += s[x]
	}
	return r
}

type FSlice []float32

func (s FSlice) Sum() any {
	r := float32(0)
	for x := len(s) - 1; x > -1; x-- {
		r += s[x]
	}
	return r
}

func main() {
	var s Summable

	s = ISlice{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", s, s, s.Sum())

	s = FSlice{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", s, s, s.Sum())
}
