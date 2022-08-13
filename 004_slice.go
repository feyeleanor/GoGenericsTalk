package main

import "fmt"

func main() {
	var s []int

	s = append(s, 0, 1, 2)
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("s[:1]: %v (%T)\n", s[:1], s[:1])
	fmt.Printf("s[1:]: %v (%T)\n", s[1:], s[1:])
	s = append(s, s...)
	fmt.Printf("s: %v (%T)\n", s, s)
	s = make([]int, 3, 6)
	fmt.Printf("s: %v (%T)\n", s, s)
	s = []int{2, 4, 6}
	s[1] = s[2]
	for i, v := range s {
		fmt.Printf("s[%v]: %v\n", i, v)
	}
}
