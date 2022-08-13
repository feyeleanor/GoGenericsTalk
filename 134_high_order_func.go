package main

import (
	"errors"
	"fmt"
)

type Numeric interface {
	~int | ~float32
}

type Iterable interface {
	Range(f any) error
}

func Sum(s Iterable) (r any) {
	switch s := s.(type) {
	case []int:
		var x int
		s.Range(func(i, v int) {
			x += v
		})
		r = x
	case []float32:
		var x float32
		for _, v := range s {
			x += v
		}
		r = x
	}
	return
}

type ISlice []int

func (s ISlice) Range(f any) (e error) {
	switch f := f.(type) {
	case func(int):
		for _, v := range s {
			f(v)
		}
	case func(int, int):
		for i, v := range s {
			f(i, v)
		}
	default:
		e = errors.New("iterator function cannot handle int parameters")
	}
	return
}

type FSlice []float32

func (s FSlice) Range(f any) (e error) {
	switch f := f.(type) {
	case func(float32):
		for _, v := range s {
			f(v)
		}
	case func(int, float32):
		for i, v := range s {
			f(i, v)
		}
	default:
		e = errors.New("iterator function cannot handle float32 parameters")
	}
	return
}

func main() {
	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v [%T]\n", is, Sum(is), is[0])

	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%v) = %v [%T]\n", fs, Sum(fs), fs[0])
}
