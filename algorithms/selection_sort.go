package main

import "fmt"

func selectSort(numbers []int) []int {

	for i := range numbers {
		min := i
		j := i + 1
		for j < len(numbers) {
			if numbers[j] < numbers[i] {
				min = j
			}
			j++
		}
		numbers[i], numbers[min] = numbers[min], numbers[i]
	}

	return numbers

}
func main() {

	values := []int{1, 3, 2, 4}
	fmt.Println(selectSort(values))

}
