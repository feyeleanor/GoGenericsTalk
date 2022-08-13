package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Summable[T Numeric] []T

func (s Summable[T]) Sum() (r T) {
	for _, x := range s {
		r += x
	}
	return
}

type ISlice = Summable[int]
type FSlice = Summable[float32]

func main() {
	i := Summable[int]{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", i, i, i.Sum())

	i = ISlice{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", i, i, i.Sum())

	f := Summable[float32]{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", f, f, f.Sum())

	f = FSlice{0, 1, 2, 3, 4}
	fmt.Printf("(%T)%v.Sum() = %v\n", f, f, f.Sum())
}
