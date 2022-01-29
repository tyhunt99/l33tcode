package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := fmt.Sprint(x)
	var r string

	// reverse the string
	for _, l := range n {
		r = string(l) + r
	}

	return n == r
}

func main() {
	fmt.Printf("isPalindrome(121) = %t\n", isPalindrome(121))
}
