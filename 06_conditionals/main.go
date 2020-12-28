package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x < y && x != 0 {
		fmt.Printf("%d is smaller than %d\n", x, y)
	} else {
		fmt.Printf("%d is smaller than %d\n", y, x)
	}

	color := "green"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is neither red nor blue")
	}

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is neither red nor blue")
	}
}
