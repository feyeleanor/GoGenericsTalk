package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Iterable[T Numeric] interface {
	Range(f func(int, T))
}

type Reducible[T Numeric] interface {
	Reduce(func(T, T) T) T
}

func Reduce[T Numeric](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case Iterable[T]:
		var x T
		c.Range(func(i int, v T) {
			x += v
		})
		r = x
	case Reducible[T]:
		r = c.Reduce(f)
	case string:
		for _, x := range c {
			r = f(r, T(x))
		}
	case T:
		r = f(r, c)
	case []T:
		for _, x := range c {
			r = f(r, x)
		}
	case map[int]T:
		for _, x := range c {
			r = f(r, x)
		}
	case chan T:
		for x := range c {
			r = f(r, x)
		}
	default:
		switch c := R.ValueOf(c); c.Kind() {
		case R.Map:
			for i := c.MapRange(); i.Next(); {
				r = f(r, i.Value().Interface().(T))
			}
		case R.Array:
			for i := 0; i < c.Len(); i++ {
				r = f(r, c.Index(i).Interface().(T))
			}
		}
	}
	return
}

type ISlice []int

func (s ISlice) Range(f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

type FArray [3]float32

func (a FArray) Reduce(f func(x, v float32) float32) (r float32) {
	for _, v := range a {
		r = f(r, v)
	}
	return
}

func DoReduce[T Numeric](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

func Pump[T Numeric](ic chan<- T, v ...T) {
	for _, v := range v {
		ic <- v
	}
	close(ic)
}

func Adder[T Numeric]() func(T, T) T {
	return func(x, v T) T {
		return x + v
	}
}

func main() {
	fi := Adder[int]()
	ff := Adder[float32]()

	DoReduce(ISlice{0, 1, 2, 3, 4}, fi)
	DoReduce(FArray{0, 1, 2}, ff)
	DoReduce("01234", fi)
	DoReduce(10, fi)
	DoReduce(10.0, ff)
	DoReduce([]int{0, 1, 2, 3, 4}, fi)
	DoReduce([]float32{0, 1, 2, 3, 4}, ff)
	DoReduce(map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}, fi)
	DoReduce(map[int]float32{0: 0, 1: 1, 2: 2, 3: 3, 4: 4}, ff)

	ic := make(chan int)
	go Pump(ic, 0, 1, 2, 3, 4)
	DoReduce(ic, fi)

	fc := make(chan float32)
	go Pump(fc, 0, 1, 2, 3, 4)
	DoReduce(fc, ff)

	DoReduce(map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4}, fi)
	DoReduce(map[string]float32{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4}, ff)
	DoReduce([5]int{0, 1, 2, 3, 4}, fi)
	DoReduce([5]float32{0, 1, 2, 3, 4}, ff)
}
