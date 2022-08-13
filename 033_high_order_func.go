package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(f func(int, T))
}

func Sum[T Numeric](s any) (r any) {
	switch s := s.(type) {
	case Iterable[T]:
		fmt.Printf("case Iterable[%T]\n", s)
		var x T
		s.Range(func(i int, v T) {
			x += v
		})
		r = x
	case []T:
		fmt.Printf("case %T\n", s)
		var x T
		for _, v := range s {
			x += v
		}
		r = x
	}
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
	fmt.Printf("Sum(%v) = %v [%T]\n", is, Sum[int](is), is[0])

	fs := []float32{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v [%T]\n", fs, Sum[float32](fs), fs[0])
}
