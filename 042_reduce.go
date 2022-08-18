package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~float32
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case [1]T:
		for x := range c {
			r = f(r, T(x))
		}
	case [2]T:
		for _, x := range c {
			r = f(r, T(x))
		}
	default:
		switch c := R.ValueOf(c); c.Kind() {
		case R.Array:
			for i := 0; i < c.Len(); i++ {
				r = f(r, c.Index(i).Interface().(T))
			}
		}
	}
	return
}

func DoReduce(c any) {
	ir := Reduce(c, func(x, v int) int {
		return x + v
	})
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, ir, ir)
}

func main() {
	DoReduce([0]int{})
	DoReduce([1]int{0})
	DoReduce([2]int{0, 1})
	DoReduce([5]int{0, 1, 2, 3, 4})
}
