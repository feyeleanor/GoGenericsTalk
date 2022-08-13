package main

import "fmt"

type Indexable[T any] interface {
	~[]T
}

type ISlice []int

func Append[T Indexable[E], E any](s *T, e E) {
	*s = append(*s, e)
}

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	is2 := append(make(ISlice, 0), is...)
	Append(&is2, 5)
	fmt.Printf("Append[%T](%v, %v) = %T%v\n", is, is, 5, is2, is2)

	fs := []float32{0, 1, 2, 3, 4}
	fs2 := append(make([]float32, 0), fs...)
	Append(&fs2, 6)
	fmt.Printf("Append[%T](%v, %v) = %T%v\n", fs, fs, 6, fs2, fs2)
}
