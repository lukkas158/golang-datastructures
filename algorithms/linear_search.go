package main

import (
	"errors"
	"fmt"
)

// Linear Search Pseudocode
//  for i = 1 to A.length
//  	if A[i] == v
// 			return i
//  return nil

func search(numbers []int, value int) (int, error) {
	for i := range numbers {
		if numbers[i] == value {
			return i, nil
		}
	}

	return 0, errors.New("Not found")
}

func main() {

	// index, err := search([]int{1, 2, 3}, 3)
	index, err := search([]int{1, 2, 3}, 4)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(index)
}
