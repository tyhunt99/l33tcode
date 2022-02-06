package main

func lengthOfLongestSubstring(s string) int {
	var length, maxLength int
	// var substring []rune
	var ss map[rune]int
	// for each letter loop through
	runeString := []rune(s)
	for i := range runeString {
		// short circuit if since there is not enough remaining chars to be longer
		if len(runeString)-i <= maxLength {
			return maxLength
		}

		// loop through to find first repeat char
		ss = make(map[rune]int)
		length = 0
		for j := i; j < len(runeString); j++ {
			_, repeat := ss[runeString[j]]
			if repeat {
				length = 0
				break
			} else {
				length++
				ss[runeString[j]] = i
			}
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}
