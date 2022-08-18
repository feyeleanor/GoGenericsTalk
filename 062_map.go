package main

import (
	"fmt"
	R "reflect"
)

func Map[T any](s any, f func(T) T) (r []T) {
	switch s := R.ValueOf(s); s.Kind() {
	case R.Map:
		switch s.Type().Key().Kind() {
		case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
			r = make([]T, s.Len())
			for i := s.MapRange(); i.Next(); {
				x := int(i.Key().Int())
				if x >= len(r) {
					n := make([]T, x+1)
					copy(n, r)
					r = n
				}
				r[x] = f(i.Value().Interface().(T))
			}
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
	DoMap[int](map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int8]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int32]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
}
