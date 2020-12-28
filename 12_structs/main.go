package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name, email string
	age         int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Welcome " + p.name + " (" + strconv.Itoa(p.age) + ")"
}

// birthday (pointer receiver)
func (p *Person) birthday() {
	p.age++
}

func main() {
	p1 := Person{name: "Alice", age: 24}

	p2 := Person{"Bob", "bob@mail.com", 25}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.name)

	p1.age++
	fmt.Println(p1)

	p1.birthday()

	fmt.Println(p1.greet())
}
