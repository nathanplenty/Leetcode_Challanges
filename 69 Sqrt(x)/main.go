package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", mySqrt(4), "\nExpected: 2")
	fmt.Println("Output 2:", mySqrt(8), "\nExpected: 2")

	fmt.Println("STOP")
}

// Function to compute the integer square root of a non-negative integer x.
func mySqrt(x int) int {
	// Handle edge cases where x is 0 or 1.
	if x == 0 || x == 1 {
		return x
	}

	// Initialize variables for binary search.
	left, right := 0, x
	var result int

	// Perform binary search to find the integer square root.
	for left <= right {
		// Calculate the middle point.
		mid := left + (right-left)/2

		// Check if mid is the square root of x.
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			// If mid*mid < x, search in the right half.
			left = mid + 1
			result = mid // Keep track of the potential result as the largest integer whose square <= x.
		} else {
			// If mid*mid > x, search in the left half.
			right = mid - 1
		}
	}

	// Return the largest integer whose square <= x.
	return result
}
