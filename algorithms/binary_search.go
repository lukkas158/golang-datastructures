package main

import (
	"errors"
	"fmt"
)

func binarySearch(number int, numbers []int) (int, error) {

	lb := 0
	ub := len(numbers)

	for {
		middle := int((lb + ub) / 2.0)

		if numbers[middle] == number {
			return middle, nil
		} else if number < numbers[middle] {
			ub = middle
		} else {
			lb = middle
		}

		if middle == lb {
			break
		}

	}
	return 0, errors.New("Number not found")
}

func recursiveBinarySearch(number int, numbers []int, lb int, ub int) (int, error) {

	middle := int((lb + ub) / 2.0)
	if numbers[middle] == number {
		return middle, nil
	} else if middle == lb {
		return 0, errors.New("Number not found")
	} else if number < numbers[middle] {
		ub = middle
		return recursiveBinarySearch(number, numbers, lb, middle)
	} else {
		return recursiveBinarySearch(number, numbers, middle, ub)
	}

}

func main() {

	fmt.Println(binarySearch(0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(binarySearch(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(binarySearch(11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println("\nRecursive version")
	fmt.Println(recursiveBinarySearch(0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 9))
	fmt.Println(recursiveBinarySearch(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 9))
	fmt.Println(recursiveBinarySearch(11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, 9))

}
