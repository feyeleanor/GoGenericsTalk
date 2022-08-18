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
	case chan T:
		for x := range c {
			r = f(r, x)
		}
	case string:
		for _, x := range c {
			r = f(r, T(x))
		}
	default:
		switch c := R.ValueOf(c); c.Kind() {
		case R.Map:
			for i := c.MapRange(); i.Next(); {
				r = f(r, i.Value().Interface().(T))
			}
		}
	}
	return
}

func Pump[T Numeric](ic chan<- T, v ...T) {
	for _, v := range v {
		ic <- v
	}
	close(ic)
}

func DoReduce(c any) {
	ir := Reduce(c, func(x, v int) int {
		return x + v
	})
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, ir, ir)
}

func main() {
	ic := make(chan int)
	go Pump(ic, 0, 1, 2, 3, 4)
	DoReduce(ic)
	DoReduce("01234")
	DoReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoReduce(map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4})
}
