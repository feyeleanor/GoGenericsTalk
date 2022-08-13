package main

import "fmt"

func main() {
	var a [3]int

	fmt.Println("len(a):", len(a))
	fmt.Println("cap(a):", cap(a))
	fmt.Printf("a: %v (%T)\n", a, a)
	fmt.Printf("a[1]: %v (%T)\n", a[1], a[1])
	fmt.Printf("a[0:1]: %v (%T)\n", a[0:1], a[0:1])

	a = [3]int{3, 2, 1}
	fmt.Printf("a: %v (%T)\n", a, a)
	a[0] = a[1]
	for i, v := range a {
		fmt.Printf("a[%v]: %v\n", i, v)
	}
}
