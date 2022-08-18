package main

import (
	"fmt"
	R "reflect"
)

func Reduce[T any](c any, f func(T, T) T) (r T) {
	switch c := R.ValueOf(c); c.Kind() {
	case R.String:
		for _, x := range c.Interface().(string) {
			r = f(r, any(x).(T))
		}
	}
	return
}

func DoReduce[T any](c any, f func(T, T) T) {
	r := Reduce(c, f)
	fmt.Printf("[%T]Reduce(%v, f()) = %v[%T]\n", c, c, r, r)
}

func main() {
	DoReduce("01234", func(x, v rune) rune {
		return x + v - 48
	})
}
