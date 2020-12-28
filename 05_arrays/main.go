package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	fmt.Println(fruitArr)

	vegArr := [2]string{"tomato", "brinjal"}
	fmt.Println(vegArr)

	fruitSlice := []string{"apple", "orange", "grape", "cherry"}
	fmt.Println(fruitSlice)

	fmt.Println(len(fruitSlice)) // 4

	fmt.Println(fruitSlice[0:2]) // [apple orange]
}
