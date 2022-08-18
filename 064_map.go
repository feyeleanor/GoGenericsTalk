package main

import "fmt"

func Map[T any](s any, f func(T) T) (r []T) {
	switch s := s.(type) {
	case string:
		for _, x := range s {
			r = append(r, f(any(x).(T)))
		}
	case chan T:
		for x := range s {
			r = append(r, f(x))
		}
	}
	return
}

func DoMap[T ~int | ~rune](s any) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

func Pump[T any](ic chan<- T, v ...T) {
	for _, v := range v {
		ic <- v
	}
	close(ic)
}

func main() {
	ic := make(chan int)
	go Pump(ic, 0, 1, 2, 3, 4)

	DoMap[int](ic)
	DoMap[rune]("01234")
}
