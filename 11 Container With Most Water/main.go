package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), "\nExpected: 49")
	fmt.Println("Output 2:", maxArea([]int{1, 1}), "\nExpected: 1")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. n == height.length
	b. 2 <= n <= 10^5
	c. 0 <= height[i] <= 10^4
*/

// maxArea O(n) computes the maximum area of water that can be contained by the input heights.
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxVolume := 0

	for left < right {
		width := right - left
		minHeight := height[left]
		if height[right] < minHeight {
			minHeight = height[right]
		}
		currentVolume := minHeight * width

		if currentVolume > maxVolume {
			maxVolume = currentVolume
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolume
}

/*
Explanation of the function maxArea:

1. Initialize two pointers: `left` starting from the beginning of the array and `right` from the end.
2. Initialize `maxVolume` to keep track of the maximum water area found.
3. Use a loop to iterate until the two pointers meet.
	- Calculate the width between the two pointers.
	- Determine the shorter height between `left` and `right`.
	- Compute the current volume of water that can be contained.
	- Update `maxVolume` if the current volume is greater.
	- Move the pointer pointing to the shorter height inward, as this might help in finding a larger area.
4. Return the maximum volume found.
*/
