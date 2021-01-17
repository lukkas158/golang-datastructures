package main

import "fmt"

type Vertex struct {
	Lat, Lng float64
}

func main() {

	var places = make(map[string]Vertex)

	places["Bell"] = Vertex{15.3, 32.3}

	places["Bell2"] = Vertex{15.3, 32.3}

	elem, ok := places["bell3"]

	fmt.Println(elem, ok) // Check if an element is present in the map.
	fmt.Println(places)
}
