package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var threesums [][]int
	n := len(nums)

	// sort the array for som optimizations
	// allows easy duplicate removal
	// allows reduced checks because of changing values is predicatable,
	// eg move up index for increase sum or down index for decrease sum
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		// loop through 2sums
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				threesums = append(threesums, []int{nums[i], nums[j], nums[k]})
				j++
				// skip dupes
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > 0 {
				// decrease sum
				k--
			} else {
				// increase sum
				j++
			}

		}
	}
	return threesums
}

func main() {
	tests := []struct {
		description string
		input       []int
		expected    [][]int
	}{
		{
			description: "Basic case",
			input:       []int{-1, 0, 1, 2, -1, -4},
			expected:    [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			description: "Another case",
			input:       []int{},
			expected:    [][]int{},
		},
		{
			description: "Yet another case",
			input:       []int{1, 1},
			expected:    [][]int{},
		},
		{
			description: "Duplicate case",
			input:       []int{-1, 0, 1, 2, -1, -4, 0, 0},
			expected:    [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			description: "Avoid dupes case",
			input:       []int{0, 0, 0, 0},
			expected:    [][]int{{0, 0, 0}},
		},
	}
	for _, test := range tests {
		result := threeSum(test.input)
		fmt.Println(result)
	}
}
