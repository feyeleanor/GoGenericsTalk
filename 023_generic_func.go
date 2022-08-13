package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

func Sum[T Numeric](s []T) (r T) {
	for _, x := range s {
		r += x
	}
	return
}

type ISlice []int
type FSlice []float32

func main() {
	i := []int{0, 1, 2, 3, 4}
	fmt.Printf("Sum[%T](%v) = %v\n", i[0], i, Sum(i))

	f := []float32{0, 1, 2, 3, 4}
	fmt.Printf("Sum[%T](%v) = %v\n", f[0], f, Sum(f))

	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum[%T](%v) = %v\n", is[0], is, Sum(is))

	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum[%T](%v) = %v\n", fs[0], fs, Sum(fs))
}
