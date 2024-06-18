package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", removeDuplicates([]int{1, 1, 2}), "\nExpected: 2")
	fmt.Println("Output 2:", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}), "\nExpected: 5")

	fmt.Println("STOP")
}

// Function removeDuplicates removes duplicates from the given slice nums
// and returns the number of unique elements.
func removeDuplicates(nums []int) int {
	// Initialize current with the first element of nums
	current := nums[0]
	// Initialize count to track the number of unique elements
	count := 1
	// Iterate over the slice starting from index 1
	for i := 1; i < len(nums); {
		// If current element is equal to nums[i], it's a duplicate
		if current == nums[i] {
			// Remove the duplicate element from nums
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			// Update current to nums[i], indicating a new unique element
			current = nums[i]
			// Increment count for the new unique element
			count++
			// Move to the next element in the slice
			i++
		}
	}
	// Return the number of unique elements
	return count
}
