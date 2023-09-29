package main

import (
	"fmt"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func check_grid(grid [9][9]int) bool {
	// Row Check
	cur_row := []int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if contains(cur_row, grid[i][j]) {
				return false
			} else {
				cur_row = append(cur_row, grid[i][j])
			}
		}
		cur_row = cur_row[:0]
	}
	// Column Check
	cur_col := []int{}
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if contains(cur_col, grid[i][j]) {
				return false
			} else {
				cur_col = append(cur_col, grid[i][j])
			}
		}
		cur_col = cur_col[:0]
	}
	// Grid Check ( %3)
	cur_quad := []int{}
	for qx := 0; qx < 9; qx += 3 {
		for qy := 0; qy < 9; qy += 3 {
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if contains(cur_quad, grid[qx+x][qy+y]) {
						return false
					} else {
						cur_quad = append(cur_quad, grid[qx+x][qy+y])
					}
				}
			}
			cur_quad = cur_quad[:0]
		}
	}
	return true
}

func main() {
	// make good grid
	board := [9][9]int{{3, 1, 6, 5, 7, 8, 4, 9, 2},
		{5, 2, 9, 1, 3, 4, 7, 6, 8},
		{4, 8, 7, 6, 2, 9, 5, 3, 1},
		{2, 6, 3, 4, 1, 5, 9, 8, 7},
		{9, 7, 4, 8, 6, 3, 1, 2, 5},
		{8, 5, 1, 7, 9, 2, 6, 4, 3},
		{1, 3, 8, 9, 4, 7, 2, 5, 6},
		{6, 9, 2, 3, 5, 1, 8, 7, 4},
		{7, 4, 5, 2, 8, 6, 3, 1, 9}}
	if check_grid(board) {
		fmt.Println("Good")
	} else {
		fmt.Println("Bad")
	}
}
