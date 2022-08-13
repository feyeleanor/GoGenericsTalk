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

func Map[T Numeric](s any, f func(T) T) (r []T) {
	switch s := s.(type) {
	case Iterable[T]:
		s.Range(func(i int, v T) {
			r = append(r, f(v))
		})
	case string:
		for _, x := range s {
			r = append(r, f(T(x)))
		}
	case []T:
		for _, v := range s {
			r = append(r, f(v))
		}
	case map[int]T:
		r = make([]T, len(s))
		for i, v := range s {
			if i >= len(r) {
				n := make([]T, i-1)
				copy(n, r)
				r = n
			}
			r[i] = f(v)
		}
	default:
		switch s := R.ValueOf(s); s.Kind() {
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
		case R.Func:
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
	}
	return
}

func DoMap[T Numeric](s any) {
	r := Map(s, func(v T) T {
		return v * 2
	})
	fmt.Printf("Map(%v, func()): %v\n", s, r)
}

type ISlice []int

func (s ISlice) Range(f func(i, v int)) {
	for i, v := range s {
		f(i, v)
	}
}

type NFunc[T Numeric] func(int) (T, bool)

func (f NFunc[T]) Range(p func(i int, v T)) {
	for i := 0; ; i++ {
		if r, ok := f(i); ok {
			p(i, r)
		} else {
			break
		}
	}
}

func Limit[T Numeric](i, j int, f NFunc[T]) NFunc[T] {
	return func(x int) (r T, ok bool) {
		if i <= x && x <= j {
			r, ok = f(x)
		}
		return
	}
}

func main() {
	DoMap[int](ISlice{0, 1, 2, 3, 4})
	DoMap[int]([]int{0, 1, 2, 3, 4})
	DoMap[float32]([]float32{0, 1, 2, 3, 4})

	DoMap[int]([3]int{0, 1, 2})
	DoMap[int](map[int]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int8]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})
	DoMap[int](map[int32]int{0: 0, 1: 1, 2: 2, 3: 3, 4: 4})

	DoMap[int](Limit(0, 4, func(x int) (int, bool) {
		return x, true
	}))
	DoMap[float32](func(x int) (float32, bool) {
		return float32(x), (x < 5)
	})
	DoMap[float32](func(x int) (float64, bool) {
		return float64(x), (x < 5)
	})
}
