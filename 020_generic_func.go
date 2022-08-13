package main

import (
	"fmt"
	"reflect"
)

type ISlice []int
type FSlice []float32

func Sum(s any) (r any) {
	if s := reflect.ValueOf(s); s.Kind() == reflect.Slice {
		switch s.Type().Elem().Kind() {
		case reflect.Int:
			var x int
			for i := 0; i < s.Len(); i++ {
				x += int(s.Index(i).Int())
			}
			r = x
		case reflect.Float32:
			var x float32
			for i := 0; i < s.Len(); i++ {
				x += float32(s.Index(i).Float())
			}
			r = x
		}
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
