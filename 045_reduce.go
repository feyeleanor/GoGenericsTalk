package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(f func(int, T))
}

type Reducible[T Numeric] interface {
	Reduce(func(T, T) T) T
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case func(int) (T, bool):
		for i := 0; ; i++ {
			if v, ok := c(i); ok {
				r = f(r, v)
			} else {
				break
			}
		}
	}
	return
}

func DoReduce[T Numeric](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

func Adder[T func(I) (E, bool), I, E Numeric](f T) func(E, E) E {
	return func(x, v E) E {
		return x + v
	}
}

func main() {
	i := func(x int) (int, bool) {
		return x, (x < 5)
	}
	DoReduce(i, Adder(i))

	f := func(x int) (float32, bool) {
		return float32(x), (x < 5)
	}
	DoReduce(f, Adder(f))
}
