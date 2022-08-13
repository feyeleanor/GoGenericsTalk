package main

import (
	"fmt"
	R "reflect"
)

type Numeric interface {
	~int | ~float32
}

func Map[T Numeric](s any, f func(T) T) (r []T) {
	switch s := R.ValueOf(s); s.Kind() {
	case R.Map:
		switch s.Type().Key().Kind() {
		case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
			r = make([]T, s.Len())
			for i := s.MapRange(); i.Next(); {
				x := int(i.Key().Int())
				if x >= len(r) {
					n := make([]T, x+1)
					copy(n, r)
					r = n
				}
				r[x] = f(i.Value().Interface().(T))
			}
		}
	case R.Array:
		t := s.Type().Elem()
		v := R.New(t)
		e := v.Elem()
		for i := 0; i < s.Len(); i++ {
			switch t.Kind() {
			case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
				e.SetInt(int64(s.Index(i).Int()))
				r = append(r, f(T(e.Int())))
			case R.Float32, R.Float64:
				e.SetFloat(float64(s.Index(i).Int()))
				r = append(r, f(T(e.Float())))
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
	DoMap[int](map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int8]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int32]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int]([3]int8{0, 1, 2})
}
