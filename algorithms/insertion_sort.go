package main

import "fmt"

func insertSort(numbers []int) []int {

	for j := 1; j < len(numbers); j++ {
		numberToInsert := numbers[j]
		i := j - 1 // Subarray bounds.

		//  I Must be greater than zero and must be smaller than number before it. We then can move it to the left.
		// Change the inequality to get it in descending order
		for i >= 0 && numbers[i] > numberToInsert {
			numbers[i+1] = numbers[i]
			i--
		}

		// Must be i+1 because you decrement it in the loop to prevent an infinity loop.
		numbers[i+1] = numberToInsert
	}

	return numbers
}

func main() {

	// sorted := insertSort([]int{5, 2, 4, 6, 1, 3})
	sorted := insertSort([]int{58, 31, 41, 59, 26, 41})

	for _, value := range sorted {
		fmt.Printf("%v ", value)
	}
}
