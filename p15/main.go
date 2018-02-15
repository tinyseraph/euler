package main

import (
	"fmt"
)

func main() {

	height_block := 20
	width_block := 20

	grids := make([][]int, height_block+1)
	for i := range grids {
		grids[i] = make([]int, width_block+1)
	}

	for i := range grids[0] {
		grids[0][i] = 1
	}
	for i := range grids {
		grids[i][0] = 1
	}
	fmt.Printf("%v \n", grids)

	for i := 1; i < height_block+1; i += 1 {
		for j := 1; j < width_block+1; j += 1 {
			grids[i][j] = grids[i-1][j] + grids[i][j-1]
		}
	}
	fmt.Printf("%v \n", grids)
	fmt.Printf("route: %v \n", grids[height_block][width_block])

}
