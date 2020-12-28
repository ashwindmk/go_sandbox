package main

import "fmt"

func main() {
	// Array
	ids := []int{12, 14, 24, 25, 26}

	for i, id := range ids {
		fmt.Printf("%d ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Map
	emailMap := map[string]string{"Alice": "alice@yahoo.com", "Bob": "bob@gmail.com"}

	for k, v := range emailMap {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emailMap {
		fmt.Printf("Name: %s\n", k)
	}
}
