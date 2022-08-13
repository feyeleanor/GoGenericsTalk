package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println("len(s):", len(s))
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("s[0:1]: %v (%T)\n", s[0:1], s[0:1])
	s = "cba"
	for i, v := range s {
		fmt.Printf("s[%v]: %v (%T)\n", i, v, v)
	}
	b := []byte(s)
	for i, v := range b {
		fmt.Printf("b[%v]: %v (%T)\n", i, v, v)
	}
}
