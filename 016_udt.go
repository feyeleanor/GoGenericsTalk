package main

import "fmt"

type ISlice []int

func (s ISlice) Sum() (r int) {
	for _, x := range s {
		r += x
	}
	return
}

func main() {
	s := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("%v.Sum() = %v\n", s, s.Sum())
}
