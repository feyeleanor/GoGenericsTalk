package main

import "fmt"

func main() {
	var s string

	fmt.Println("len(s):", len(s))
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("s[0:1]: %v (%T)\n", s[0:1], s[0:1])
}
