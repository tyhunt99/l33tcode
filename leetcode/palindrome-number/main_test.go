package main

import "testing"

func TestAddNumber(t *testing.T) {
	tests := []struct {
		description string
		input       int
		expected    bool
	}{
		{
			description: "Base case true",
			input:       121,
			expected:    true,
		},
		{
			description: "Base case false",
			input:       122,
			expected:    false,
		},
		{
			description: "Negative",
			input:       -121,
			expected:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if isPalindrome(test.input) != test.expected {
				t.Errorf("Bad value for %d, wanted %t", test.input, test.expected)
			}
		})
	}
}
