package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

func Map[T Numeric](s []T, f func(T) T) (r []T) {
	for _, v := range s {
		r = append(r, f(v))
	}
	return
}

func DoMap[T Numeric](s []T) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

func main() {
	DoMap([]int{0, 1, 2, 3, 4})
	DoMap([]float32{0, 1, 2, 3, 4})
}
