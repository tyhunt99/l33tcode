package addtwo

import "testing"

func buildLinkedList(num []int) *ListNode {
	var node, firstNode, previousNode *ListNode
	for i := 0; i < len(num); i++ {
		node = &ListNode{num[i], nil}
		if i == 0 {
			firstNode = node
		} else {
			previousNode.Next = node
		}
		previousNode = node
	}
	return firstNode
}

func TestAddNumber(t *testing.T) {
	tests := []struct {
		description string
		l1          []int
		l2          []int
		expected    string
	}{
		{
			description: "Example numbers",
			l1:          []int{2, 4, 3},
			l2:          []int{5, 6, 4},
			expected:    "807",
		},
		{
			description: "Small numbers",
			l1:          []int{1, 1},
			l2:          []int{1, 2},
			expected:    "32",
		},
		{
			description: "Extra carry at end",
			l1:          []int{9, 9, 9, 9, 9, 9, 9},
			l2:          []int{9, 9, 9, 9},
			expected:    "10009998",
		},
		{
			description: "Different length numbers",
			l1:          []int{1, 1},
			l2:          []int{1, 1, 2},
			expected:    "222",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := addTwoNumbers(buildLinkedList(test.l1), buildLinkedList(test.l2)).printNumber()
			if result != test.expected {
				t.Errorf("Bad number, got %s wanted %s", result, test.expected)
			}
		})
	}
}
