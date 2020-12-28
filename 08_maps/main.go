package main

import "fmt"

func main() {
	wordCount := make(map[string]int)

	arr := []string{"pen", "apple", "pen", "pine"}

	for i, s := range arr {
		fmt.Println(i, s)
		wordCount[s] = wordCount[s] + 1
	}

	fmt.Println(wordCount)

	fmt.Println(len(wordCount))

	delete(wordCount, "pine")
	fmt.Println(wordCount)

	// Direct assignment
	wordMap := map[string]int{"pen": 2, "apple": 1, "pine": 1}
	fmt.Println(wordMap)
}
