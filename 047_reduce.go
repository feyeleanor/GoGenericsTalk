package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~float32
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	if c := R.ValueOf(c); c.Kind() == R.Func {
		for i := 0; ; i++ {
			p := []R.Value{R.ValueOf(i)}
			if p = c.Call(p); p[1].Interface() == true {
				r = f(r, p[0].Interface().(T))
			} else {
				break
			}
		}
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
