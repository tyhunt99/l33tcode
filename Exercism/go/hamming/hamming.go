// Package hamming calculates the differences betweeen 2 stands of DNA
// to deteermine the hamming distance
package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between 2 DNA strings
// assumes same length strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be same length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
