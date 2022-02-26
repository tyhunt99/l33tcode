package main

import (
	"fmt"
	"testing"
)

func TestAddNumber(t *testing.T) {
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
		// {
		// 	description: "Another case",
		// 	input:       []int{},
		// 	expected:    [][]int{},
		// },
		// {
		// 	description: "Yet another case",
		// 	input:       []int{1, 1},
		// 	expected:    [][]int{},
		// },
		// {
		// 	description: "Duplicate case",
		// 	input:       []int{-1, 0, 1, 2, -1, -4, 0, 0},
		// 	expected:    [][]int{{-1, -1, 2}, {-1, 0, 1}},
		// },
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := threeSum(test.input)
			fmt.Printf("%v", result)
		})
	}
}
