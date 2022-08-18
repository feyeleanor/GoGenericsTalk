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
		c.Range(func(i int, v T) {
			r = f(r, v)
		})
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
	case func(int) (T, bool):
		for i := 0; ; i++ {
			if v, ok := c(i); ok {
				r = f(r, v)
			} else {
				break
			}
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
		case R.Func:
			for i := 0; ; i++ {
				p := []R.Value{R.ValueOf(i)}
				if p = c.Call(p); p[1].Interface() == true {
					r = f(r, p[0].Interface().(T))
				} else {
					break
				}
			}
		}
	}
	return
}
