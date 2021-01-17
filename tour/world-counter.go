package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var result = make(map[string]int)
	words := strings.Fields(s)
	// Iterate all the words of the phrase.
	for _, word := range words {
		result[word] = 0

		// Counts how many times has the word been found in the phrase.
		for _, slice := range words {
			if slice == word {
				result[word]++
			}
		}
	}

	fmt.Println("My result:", result)
	return result
}

func main() {
	wc.Test(WordCount)
}
