package main

import "fmt"

type Numeric interface {
	~int | ~float32
}

type Iterable[T Numeric] interface {
	Range(func(int, T))
}

func Map[T Numeric](s any, f func(T) T) (r []T) {
	switch s := s.(type) {
	case NFunc[T]:
		for i := 0; ; i++ {
			if v, ok := s(i); ok {
				r = append(r, f(v))
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
	DoMap[int](Limit(0, 4, func(x int) (int, bool) {
		return x, true
	}))
	DoMap[float32](func(x int) (float32, bool) {
		return float32(x), (x < 5)
	})
}
