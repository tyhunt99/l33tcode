package main

import "testing"

func TestAddNumber(t *testing.T) {
	tests := []struct {
		description string
		input       []int
		expected    int
	}{
		{
			description: "Basic case",
			input:       []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected:    49,
		},
		{
			description: "Another case",
			input:       []int{3, 9, 3, 4, 7, 2, 12, 6},
			expected:    45,
		},
		{
			description: "Yet another case",
			input:       []int{1, 1},
			expected:    1,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := maxArea(test.input)
			if result != test.expected {
				t.Errorf("Bad value for %v, got %d wanted %d", test.input, result, test.expected)
			}
		})
	}
}
