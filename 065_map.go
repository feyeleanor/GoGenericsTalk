package main

import R "reflect"

func Map[T any](s any, f func(T) T) (r []T) {
	switch s := s.(type) {
	case interface{ Range(f func(int, T)) }:
		s.Range(func(i int, v T) {
			r = append(r, f(v))
		})
	case string:
		for _, x := range s {
			r = append(r, f(any(x).(T)))
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
	case chan T:
		for x := range s {
			r = append(r, f(x))
		}
	default:
		switch s := R.ValueOf(s); s.Kind() {
		case R.Array:
			for i := 0; i < s.Len(); i++ {
				r = append(r, f(s.Index(i).Interface().(T)))
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
			for i := 0; ; i++ {
				p := []R.Value{R.ValueOf(i)}
				if p = s.Call(p); p[1].Interface() == true {
					r = append(r, f(p[0].Interface().(T)))
				} else {
					break
				}
			}
		}
	}
	return
}

func main() {}
