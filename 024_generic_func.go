package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Indexable[T any] interface {
	~[]T
}

func Sum[T Indexable[E], E Numeric](s T) (r E) {
	for _, x := range s {
		r += x
	}
	return
}

type ISlice Indexable[int]

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum[%T](%v) = %v\n", is[0], is, Sum(is))
}
