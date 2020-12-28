package main

import "fmt"

func main() {
	// Method 1
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Method 2
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
