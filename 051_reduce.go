package main

import "fmt"

func Reduce[T any](c any, f func(T, T) T) (r T) {
	switch c := c.(type) {
	case string:
		for _, x := range c {
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
	DoReduce[rune]("01234", func(r, x rune) rune {
		return r + x - 48
	})
}
