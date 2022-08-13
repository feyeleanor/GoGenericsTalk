package main

import "fmt"

func main() {
	var s []int

	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
	fmt.Printf("s: %v (%T)\n", s, s)
	fmt.Printf("s[0:1]: %v (%T)\n", s[0:1], s[0:1])
}
