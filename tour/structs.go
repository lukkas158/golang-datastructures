package main

import "fmt"

// Declaring a Struct
type Vertex struct {
	x int
	Y int
}

func main() {
	// Instacing a Struct
	vertex := Vertex{1, 2}
	vertexNamed := Vertex{x: 10}

	fmt.Println("Vertex Named", vertexNamed)

	vertex2 := vertex    // Creates a copy
	fmt.Println(vertex2) // {1 12}
	vertex2.x = 10

	fmt.Println(vertex2)  // {10, 2}
	fmt.Println(vertex)   // {1 2}
	fmt.Println(vertex.x) // 1

	vertex3 := &vertex // Pointers
	vertex3.x = 10     // Mutates vertex

	fmt.Println(vertex.x) // 10

}
