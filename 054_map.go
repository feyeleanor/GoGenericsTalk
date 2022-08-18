package main

import "fmt"

func Map[T any](s []T, f func(T) T) (r []T) {
	for _, v := range s {
		r = append(r, f(v))
	}
	return
}

func DoMap[T ~int | ~float32](s []T) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

func main() {
	DoMap([]int{0, 1, 2, 3, 4})
	DoMap([]float32{0, 1, 2, 3, 4})
}
