package main

import "fmt"

func Reduce(s any, f func(any, any) any) (r any) {
	switch s := s.(type) {
	case []int:
		var x int
		for _, v := range s {
			x = f(x, v).(int)
		}
		r = x
	case []float32:
		var x float32
		for _, v := range s {
			x = f(x, v).(float32)
		}
		r = x
	}
	return
}

func main() {
	is := []int{0, 1, 2, 3, 4}
	ir := Reduce(is, func(x any, v any) any {
		return x.(int) + v.(int)
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", is, ir, ir)

	fs := []float32{0, 1, 2, 3, 4}
	fr := Reduce(fs, func(x any, v any) any {
		return x.(float32) + v.(float32)
	})
	fmt.Printf("Reduce(%v, f()) = %v [%T]\n", fs, fr, fr)
}
