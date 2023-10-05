package main

import (
	"fmt"
)

func volume(x int, y1, y2 int) int {
	if y1 < y2 {
		return x * y1
	} else {
		return x * y2
	}
}

func process_array(array []int) int {
	max := 0
	for index, val := range array {
		for index2, val2 := range array {
			// calculate delta-x
			delta_x := index2 - index
			if delta_x <= 0 {
				continue
			}
			area := volume(delta_x, val, val2)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func main() {
	// make array (slice) here
	made_array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Array: ", made_array)
	result := process_array(made_array)
	fmt.Println(result)
}
