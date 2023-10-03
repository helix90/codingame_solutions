package main

import (
	"fmt"
	"strings"
)

func max_height(mountains []int64) int {
	// find the highest mountain
	max_height := 0
	for _, val := range mountains {
		if int(val) > max_height {
			max_height = int(val)
		}
	}
	return max_height
}

func print_mountains(mountains []string) {
	// Print the entries in the map of the mountains
	for _, val := range mountains {
		// Remember to trim the trailing spaces
		fmt.Println(strings.TrimRight(val, " "))
	}
}

func build_output(mountains []int64) []string {
	// Take the map of the heights, and make the printable strings
	max_height := max_height(mountains)
	ret_val := []string{}
	var counter int
	// Layer must be external to the loop for scope reasons
	// Do no use a MAP, use a Slice.
	// Maps are not ordered, and order matters here.
	var layer []string
	for altitude := max_height; altitude > 0; altitude-- {
		// Altitude of this pass, from highest to lowest
		counter = 0
		for _, val := range mountains {
			layer = layer[:0]
			for i := 0; i < 2*int(val); i++ {
				layer = append(layer, ` `)
			}
			if (int(val) - altitude) >= 0 {
				left := altitude - 1
				right := 2*int(val) - altitude
				layer[left] = `/`
				layer[right] = `\`
			}
			layer_string := strings.Join(layer, "")
			if counter == 0 {
				ret_val = append(ret_val, layer_string)
			} else {
				ret_val[counter-1] += layer_string
			}
			counter = max_height - altitude + 1
		}
		// Join layer into string
	}
	// Join the characters in layer
	return ret_val
}

func main() {
	mountains := []int64{}
	mountains = append(mountains, int64(2))
	mountains = append(mountains, int64(3))
	mountains = append(mountains, int64(2))
	mountains = append(mountains, int64(5))
	mountains = append(mountains, int64(1))
	output := build_output(mountains)
	print_mountains(output)
}
