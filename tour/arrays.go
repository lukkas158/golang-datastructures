package main

import "fmt"

func main() {

	// Declaring Arrays
	a := []int{1, 3, 3, 5, 6}
	var b [5]int
	var c [2]string

	b[0] = 1
	c[1] = "Lucas"

	// Slicing
	s := a[0:3]

	fmt.Println(a, b, c, s)
	// Be Careful, slice is a pointer to the orignal array;
	s[1] = 1000

	fmt.Println(a)

	// Creating an Array of Structs

	vertexs := []struct {
		x int
		y int
	}{
		{2, 10},
		{1, 5},
		{5, 5},
	}

	appenedVertex := append(vertexs, struct{ x, y int }{10, 20})

	fmt.Println("Appended", appenedVertex)
	fmt.Println(vertexs)

	var nilSlice []int

	if nilSlice == nil {
		println("An empty slice with capacity zero is equal to nil")
	}

}
