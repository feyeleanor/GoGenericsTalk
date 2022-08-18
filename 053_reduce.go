package main

import R "reflect"

func Reduce[T any](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case interface{ Range(f func(int, T)) }:
		c.Range(func(i int, v T) {
			r = f(r, v)
		})
	case interface{ Reduce(func(T, T) T) T }:
		r = c.Reduce(f)
	case string:
		for _, x := range c {
			r = f(r, any(x).(T))
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

func main() {}
