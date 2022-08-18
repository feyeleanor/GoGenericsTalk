package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Reducible[K comparable, E Numeric] map[K]E

func (m Reducible[K, E]) Reduce(f func(E, E) E) (r E) {
	for _, x := range m {
		r = f(r, x)
	}
	return
}

type IMap = Reducible[int, int]

func main() {
	im := IMap{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}
	ir := im.Reduce(func(x, v int) int {
		return x + v
	})
	fmt.Printf("%v.Reduce(f()) = %v [%T]\n", im, ir, ir)

	fm := Reducible[int, float32]{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}
	fr := fm.Reduce(func(x, v float32) float32 {
		return x + v
	})
	fmt.Printf("%v.Reduce(f()) = %v [%T]\n", fm, fr, fr)
}
