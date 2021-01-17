package main

import "fmt"

func main() {
	// Go Lang has pointers almost equal to C. Unlike C, it has not pointer arithmetic.
	i := 10

	p := &i // Set p to pointer to i

	fmt.Println(*p) // Reads i through p.

	*p = *p / 2 // Write to i through p

	fmt.Println(i)

}
