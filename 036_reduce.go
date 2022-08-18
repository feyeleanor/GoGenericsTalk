package main

import (
	"fmt"
	R "reflect"
)

func Reduce(s any, f func(any, any) any) (r any) {
	if s := R.ValueOf(s); s.Kind() == R.Slice {
		T := s.Type().Elem()
		V := R.New(T)
		E := V.Elem()
		switch T.Kind() {
		case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
			for i := 0; i < s.Len(); i++ {
				v := R.ValueOf(f(E.Interface(), s.Index(i).Interface()))
				E.SetInt(v.Int())
			}
		case R.Float32, R.Float64:
			for i := 0; i < s.Len(); i++ {
				v := R.ValueOf(f(E.Interface(), s.Index(i).Interface()))
				E.SetFloat(v.Float())
			}
		}
		r = E.Interface()
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
