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
			input:       []int{7, 1, 5, 3, 6, 4},
			expected:    5,
		},
		{
			description: "No profit possible",
			input:       []int{7, 6, 4, 3, 1},
			expected:    0,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := maxProfit(test.input)
			if result != test.expected {
				t.Errorf("Bad value for %v, got %d wanted %d", test.input, result, test.expected)
			}
		})
	}
}
