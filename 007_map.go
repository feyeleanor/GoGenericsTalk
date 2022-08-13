package main

import "fmt"

func main() {
	var m map[int]int

	fmt.Println("len(m):", len(m))
	fmt.Println("m[1]:", m[1])
	m[1] = 1
}
