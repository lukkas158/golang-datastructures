package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	slices := make([][]uint8, dy)

	for i := range slices {
		slices[i] = make([]uint8, dx)
		for j := range slices[i] {
			slices[i][j] = uint8(i * j)
		}
	}

	return slices
}

func main() {
	pic.Show(Pic)
}
