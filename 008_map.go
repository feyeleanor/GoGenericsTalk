package main

import "fmt"

func main() {
	var m map[int]int

	m = make(map[int]int)
	m[1] = 1
	fmt.Println("m[1]:", m[1])
	delete(m, 1)
	fmt.Println("m:", m)
	x, ok := m[1]
	fmt.Printf("x: %v, ok: %v\n", x, ok)
	m = map[int]int{0: 2, 3: 6}
	fmt.Println("m:", m)
	for k, v := range m {
		fmt.Printf("m[%v]: %v\n", k, v)
	}
}
