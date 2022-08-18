package main

import (
	"fmt"
	R "reflect"
)

func Map[T any](s any, f func(T) T) (r []T) {
	if s := R.ValueOf(s); s.Kind() == R.Array {
		defer func() {
			recover()
		}()
		for i := 0; i < s.Len(); i++ {
			r = append(r, f(s.Index(i).Interface().(T)))
		}
	}
	return
}

func DoMap[T ~int | ~float32](s any) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

func main() {
	DoMap[int]([3]int{0, 1, 2})
	DoMap[float32]([3]float32{0, 1, 2})
	DoMap[float32]([3]float64{0, 1, 2})
}
