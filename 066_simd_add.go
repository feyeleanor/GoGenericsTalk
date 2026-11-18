//go:build goexperiment.simd

package main

import "fmt"
import "simd"

type Numeric interface {
	~int64 | ~float64
}

var vecF64 simd.Float64s
var vecI64 simd.Int64s

func vecLen[T Numeric](x []T) (r int) {
	switch any(x).(type) {
	case []float64:
		r = vecF64.Len()
	case []int64:
		r = vecI64.Len()
	}
	return
}

func padVector[T Numeric](x []T) []T {
	return append(x, make([]T, len(x) % vecLen(x))...)
}

func reallocate[T Numeric](x []T, c int) (r []T) {
	r = make([]T, c)
	copy(r, x)
	return
}

func matchCapacity[T Numeric](x, y []T) (rx, ry []T) {
	switch rx, ry = x, y; {
	case cap(x) < cap(y):
		rx = reallocate(rx, cap(y))
	case cap(y) < cap(x):
		ry = reallocate(ry, cap(x))
	}
	return
}

func Add[T Numeric](x, y []T) (r []T) {
	xv, yv := matchCapacity(x, y)
	r = make([]T, cap(xv), cap(xv))
	switch any(x).(type) {
	case []float64:
		vl := vecLen(xv)
		for i := 0; i < len(xv); i += vl {
			a := simd.LoadFloat64s(any(xv).([]float64)[i:i+vl])
			b := simd.LoadFloat64s(any(yv).([]float64)[i:i+vl])
			a.Add(b).Store(any(r).([]float64)[i:i+vl])
		}

	case []int64:
		vl := vecLen(xv)
		for i := 0; i < len(xv); i += vl {
			a := simd.LoadInt64s(any(xv).([]int64)[i:i+vl])
			b := simd.LoadInt64s(any(yv).([]int64)[i:i+vl])
			a.Add(b).Store(any(r).([]int64)[i:i+vl])
		}
	}
	return
}

func DoAdd[T Numeric](x, y []T) {
	fmt.Printf("Add(%v, %v):\n", x, y)
	r := Add(x, y)
	x, y = matchCapacity(x, y)
	for i, v := range r {
		fmt.Println("\t", x[i], "\t+", y[i], "\t=", v)
	}
}

func main() {
	DoAdd(
		[]float64{0, 1, 4, 10, 22, 46},
		[]float64{0, 2, 4, 6, 8, 10, 12, 14})
	DoAdd(
		[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		[]int64{0, 2, 4, 6, 8, 10, 12, 14})
}
