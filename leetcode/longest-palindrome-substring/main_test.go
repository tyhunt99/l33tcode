package main

import "testing"

func TestAddNumber(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "Basic case",
			input:       "babad",
			expected:    "bab",
		},
		{
			description: "Example 2",
			input:       "cbbd",
			expected:    "bb",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := longestPalindrome(test.input)
			if result != test.expected {
				t.Errorf("Bad value for %s, got %s wanted %s", test.input, result, test.expected)
			}
		})
	}
}
