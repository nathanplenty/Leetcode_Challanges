package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", searchInsert([]int{1, 3, 5, 6}, 5), "\nExpected: 2")
	fmt.Println("Output 2:", searchInsert([]int{1, 3, 5, 6}, 2), "\nExpected: 1")
	fmt.Println("Output 3:", searchInsert([]int{1, 3, 5, 6}, 7), "\nExpected: 4")

	fmt.Println("STOP")
}

// searchInsert O(n) takes in a sorted array of integers (nums) and a target integer.
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}

/*
Explanation of the function searchInsert:

1. Loop through each element of the array `nums`.
2. For each index `i`:
   a. Check if the `target` is less than or equal to the current element `nums[i]`.
   b. If yes, return the index `i`.
3. If none of the elements are greater than or equal to the `target`, return the length of the array.
*/
