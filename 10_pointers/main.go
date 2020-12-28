package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // *int

	// Read memory address
	fmt.Println(b) // 0xc00001a0a8

	// Read value from address
	fmt.Println(*b) // 5

	// Change value at address
	*b = 10
	fmt.Println(a) // 10
}
