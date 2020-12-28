package main

import "fmt"

func main() {
	fmt.Println(greeting("John"))
	fmt.Println(getSum(3, 4))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(n1, n2 int) int {
	return n1 + n2
}
