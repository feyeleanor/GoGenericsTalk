package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(func(int, T))
}

func Map[T Numeric](s any, f func(T) T) (r []T) {
	switch s := s.(type) {
	case Iterable[T]:
		s.Range(func(i int, v T) {
			r = append(r, f(v))
		})
	case []T:
		for _, v := range s {
			r = append(r, f(v))
		}
	}
	return
}

func DoMap[T Numeric](s []T) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

type ISlice []int

func (s ISlice) Range(f func(i, v int)) {
	for i, v := range s {
		f(i, v)
	}
}

func main() {
	DoMap(ISlice{0, 1, 2, 3, 4})
	DoMap([]int{0, 1, 2, 3, 4})
	DoMap([]float32{0, 1, 2, 3, 4})
}
