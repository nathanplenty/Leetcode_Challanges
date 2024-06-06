package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", searchInsert([]int{1, 3, 5, 6}, 5), "\nExpected: 2")
	fmt.Println("Output 2:", searchInsert([]int{1, 3, 5, 6}, 2), "\nExpected: 1")
	fmt.Println("Output 3:", searchInsert([]int{1, 3, 5, 6}, 7), "\nExpected: 4")

	fmt.Println("STOP")
}

// searchInsert takes in a sorted array of integers (nums) and a target integer.
func searchInsert(nums []int, target int) int {
	// Loop through each element of the array.
	for i := 0; i < len(nums); i++ {
		// Check if the target is less than or equal to the current element.
		if target <= nums[i] {
			// If yes, return the index of the current element.
			return i
		}
	}
	// If none of the elements are greater than or equal to the target, return the length of the array.
	return len(nums)
}
