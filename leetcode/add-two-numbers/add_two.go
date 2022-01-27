package addtwo

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode, previousNode *ListNode
	zeroNode := ListNode{0, nil}
	carry := 0
	for p := 0; p < 100; p++ {
		value := l1.Val + l2.Val + carry
		if value >= 10 {
			carry = 1
			value = value % 10
		} else {
			carry = 0
		}
		node := ListNode{
			Val:  value,
			Next: nil,
		}
		// be sure to grab first node for return
		if p == 0 {
			firstNode = &node
		} else {
			previousNode.Next = &node
		}
		previousNode = &node
		// found end of values so time to return
		if l1.Next == nil && l2.Next == nil {
			// but also need too handle any carry left
			if carry > 0 {
				node.Next = &ListNode{carry, nil}
			}
			return firstNode
		}
		// setup next iteration
		// use zeroNode for place holder for smaller numers
		if l1.Next == nil {
			l1 = &zeroNode
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = &zeroNode
		} else {
			l2 = l2.Next
		}
	}
	return firstNode
}

func (num *ListNode) printNumber() string {
	var ret string
	for p := 0; p < 100; p++ {
		ret = fmt.Sprint(num.Val) + ret
		if num.Next == nil {
			return ret
		} else {
			num = num.Next
		}
	}
	return ret
}

// func main() {
// 	nodeA := ListNode{
// 		1,
// 		&ListNode{
// 			1,
// 			nil,
// 		},
// 	}
// 	nodeA.printNumber()
// }
