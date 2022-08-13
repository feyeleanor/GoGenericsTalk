package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~float32
}

func Map[T Numeric](s any, f func(T) T) (r []T) {
	if s := R.ValueOf(s); s.Kind() == R.Func {
		V := R.ValueOf(r).Type().Elem()
		for i := 0; ; i++ {
			p := []R.Value{R.ValueOf(i)}
			if p = s.Call(p); p[1].Interface() == true {
				switch V.Kind() {
				case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
					r = append(r, f(T(p[0].Int())))
				case R.Float32, R.Float64:
					r = append(r, f(T(p[0].Float())))
				}
			} else {
				break
			}
		}
	}
	return
}

func DoMap[T Numeric](s any) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

func main() {
	DoMap[float32](func(x int) (float32, bool) {
		return float32(x), (x < 5)
	})
	DoMap[float32](func(x int) (float64, bool) {
		return float64(x), (x < 5)
	})
}
