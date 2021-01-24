package main

import (
	"fmt"
	"math"
)

func merge(left []int, right []int) []int {

	left_size := len(left)
	right_size := len(right)
	result_size := left_size + right_size
	result := make([]int, result_size)
	l := 0
	r := 0

	//  Merge values
	//  Stop when left or right is empty
	i := 0
	for l < left_size && r < right_size {

		if left[l] <= right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++

	}

	// Pass the ramainder to the the result.
	for l < left_size {
		result[i] = left[l]
		l++
		i++
	}
	for r < right_size {
		result[i] = right[r]
		r++
		i++
	}

	return result
}

func mergeSort(numbers []int) []int {

	size := float64(len(numbers))
	if size > 1 {
		left_index := int(math.Floor(size / 2))
		left := numbers[0:left_index]
		right := numbers[left_index:]
		a := mergeSort(left)
		b := mergeSort(right)
		return merge(a, b)
	} else {
		return numbers
	}

}

func main() {

	// Odd
	result := mergeSort([]int{32, 32, 31, 32, 12, 1212, 2121})
	fmt.Println(result)

	// Even
	result = mergeSort([]int{32, 32, 31, 1, 12, 1212})
	fmt.Println(result)

}
