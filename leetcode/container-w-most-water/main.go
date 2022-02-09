package main

func maxArea(height []int) int {
	var containerHeigth, maxVolume int
	// start with widest as good odds of being large volume due to known max
	left := 0
	right := len(height) - 1
	for left < right {
		leftHeight := height[left]
		rightHeight := height[right]
		if leftHeight < rightHeight {
			containerHeigth = leftHeight
		} else {
			containerHeigth = height[right]
		}

		width := right - left
		volume := width * containerHeigth
		if maxVolume < volume {
			maxVolume = volume
		}
		// remove the smaller since it cannot ever get bigger volume as width decreases
		if leftHeight < rightHeight {
			left++
		} else {
			right--
		}
	}
	return maxVolume
}
