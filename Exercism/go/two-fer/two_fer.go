// Package twofer is a communist tool for equal distribution
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith when given a name, returns a string with the message:
// One for ___, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
