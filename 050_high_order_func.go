package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(func(int, T))
}

func Map[T Numeric](s any, f func(T) T) (r []T) {
	if s := R.ValueOf(s); s.Kind() == R.Func {
		for i := 0; ; i++ {
			p := []R.Value{R.ValueOf(i)}
			if p = s.Call(p); p[1].Interface() == true {
				r = append(r, f(p[0].Interface().(T)))
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
	DoMap[float32](func(x int) (float64, bool) {
		return float64(x), (x < 5)
	})
}
