package main

import "fmt"

func Sum(s any) (r any) {
	switch s := s.(type) {
	case []float32:
		var f float32
		for _, x := range s {
			f += x
		}
		r = f
	case []int:
		var i int
		for _, x := range s {
			i += x
		}
		r = i
	}
	return
}

func main() {
	i := []int{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", i, i, Sum(i))

	f := []float32{0, 1, 2, 3, 4}
	fmt.Printf("Sum(%T%v) = %v\n", f, f, Sum(f))

	r := Sum(f).(float32)
	fmt.Println("r =", r)
}
