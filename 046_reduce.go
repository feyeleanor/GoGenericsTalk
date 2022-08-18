package main

import "fmt"

type Numeric interface {
	~int | ~float32
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
	case NFunc[T]:
		r = Reduce((func(int) (T, bool))(c), f)
	}
	return
}

type NFunc[T Numeric] func(int) (T, bool)

func DoReduce[T Numeric](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

func Adder[T Numeric]() func(T, T) T {
	return func(x, v T) T {
		return x + v
	}
}

func main() {
	DoReduce(func(x int) (int, bool) {
		return x, (x < 5)
	}, Adder[int]())
	DoReduce(func(x int) (float32, bool) {
		return float32(x), (x < 5)
	}, Adder[float32]())
	DoReduce(NFunc[int](func(x int) (int, bool) {
		return x, (x < 5)
	}), Adder[int]())
}
