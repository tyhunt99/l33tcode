package main

import "testing"

func TestAddNumber(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    int
	}{
		{
			description: "Basic case",
			input:       "abcabcbb",
			expected:    3,
		},
		{
			description: "Not at start",
			input:       "abcabcdef",
			expected:    6,
		},
		{
			description: "Single repeating character",
			input:       "bbbbbb",
			expected:    1,
		},
		{
			description: "Substring not sbsequence",
			input:       "pwwkew",
			expected:    3,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := lengthOfLongestSubstring(test.input)
			if result != test.expected {
				t.Errorf("Bad value for %s, got %d wanted %d", test.input, result, test.expected)
			}
		})
	}
}
