package main

import (
	"fmt"
	R "reflect"
)

type ISlice []int
type FSlice []float32

func Sum(s any) (r any) {
	if s := R.ValueOf(s); s.Kind() == R.Slice {
		T := s.Type().Elem()
		V := R.New(T)
		E := V.Elem()
		switch T.Kind() {
		case R.Int, R.Int8, R.Int16, R.Int32, R.Int64:
			for i := 0; i < s.Len(); i++ {
				E.SetInt(E.Int() + s.Index(i).Int())
			}
		case R.Float32, R.Float64:
			for i := 0; i < s.Len(); i++ {
				E.SetFloat(E.Float() + s.Index(i).Float())
			}
		}
		r = E.Interface()
	}
	return
}

func main() {
	i := []int{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", i, i, Sum(i))

	f := []float32{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", f, f, Sum(f))

	in := []interface{}{int(0), float32(1), int(2)}
	fmt.Printf("Sum(%T%v) = %v\n", in, in, Sum(in))

	is := ISlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", is, is, Sum(is))

	fs := FSlice{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", fs, fs, Sum(fs))
}
