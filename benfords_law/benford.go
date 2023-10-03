package main

import (
	"fmt"
)

func process_strings(numbers []string) []int {
	// Process the strings and return a slice of the first digits
	var return_array []int
	for _, val := range numbers {
		// Use string.Split to split on space
		// Use last element of slice
		for index := 0; index < len(val); index++ {
			// if char is number, add to slice and break
			if (val[index] < '1') || (val[index] > '9') {
				continue
			} else {
				foo := int(val[index] - '0')
				return_array = append(return_array, foo)
				break
			}
		}
	}
	return return_array
}

func benford_test(numbers []int, count int) bool {
	// return true/false based on Benford's Law
	// Make a map to count into
	counts := make(map[int]int)
	ratios := make(map[int]float32)
	sum := 0
	for _, val := range numbers {
		sum += val
		if _, ok := counts[val]; ok {
			counts[val] += 1
		} else {
			counts[val] = 1
		}
	}
	// Now compare counts to total lines
	for idx, val := range counts {
		ratios[idx] = 100.0 * (float32(val) / float32(count))
	}
	// Now test the ratios against the table
	if ratios[1] > 40.1 || ratios[1] < 20.1 {
		return true
	} else if ratios[2] > 27.6 || ratios[2] < 7.6 {
		return true
	} else if ratios[3] > 22.5 || ratios[3] < 2.5 {
		return true
	} else if ratios[4] > 19.7 {
		return true
	} else if ratios[5] > 17.9 {
		return true
	} else if ratios[6] > 16.7 {
		return true
	} else if ratios[7] > 15.8 {
		return true
	} else if ratios[8] > 15.1 {
		return true
	} else if ratios[9] > 14.6 {
		return true
	}
	return false
}

func main() {
	array := []string{"$ -10", "1", "$ 203"}
	output := process_strings(array)
	result := benford_test(output, len(array))
	fmt.Println(result)
}
