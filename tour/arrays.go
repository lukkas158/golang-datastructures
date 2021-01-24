package main

import "fmt"

func main() {

	//  String Array
	var names [2]string
	names[0] = "Lucas"
	names[1] = "Silva"

	fmt.Println(names)

	// Declaring Arrays
	a := []int{1, 3, 3, 5, 6, 7, 8, 1}
	var b [5]int // Start it with zeros.
	var c [2]string

	b[0] = 1
	c[1] = "Lucas"

	// Create array dinamically.
	values = make([]int, a[1])

	//  Append value to array.
	append(values, 20)

	// Slicing
	s := a[0:3] // [0,3)

	fmt.Println(a, b, c, s)
	// Be Careful, slice is a pointer to the orignal array;
	s[1] = 1000 // It will change the 3 value of a  to 1000

	fmt.Println(a)

	slice := []int{21, 21, 23, 32, 12, 12}
	//  The index if part of the array type. So, you can't resize it. Use slices to dinamically  change the the values.
	// Slices can change its size based on the original array.

	// Create a new Slice with capacity 4 and length 1
	slice = slice[2:3]
	fmt.Printf("Capacity: %d Length: %d \n", cap(slice), len(slice))

	// Drop the first two values of a and create the array with capacity cap(slice)  - 2 and length cap(slice) - 2
	slice = slice[:4]
	fmt.Printf("Capacity: %d Length: %d \n", cap(slice), len(slice))

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
