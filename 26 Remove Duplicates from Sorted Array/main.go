package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", removeDuplicates([]int{1, 1, 2}), "\nExpected: 2")
	fmt.Println("Output 2:", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}), "\nExpected: 5")

	fmt.Println("STOP")
}

// removeDuplicates O(n) removes duplicates from the given slice nums
// and returns the number of unique elements.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

/*
Explanation of the function removeDuplicates:

1. If the input array is empty, return 0.
2. Initialize the slow pointer at index 0.
3. Iterate with the fast pointer from index 1 to the end of the array:
   a. If nums[fast] is not equal to nums[slow], increment the slow pointer and update nums[slow] with nums[fast].
4. Return slow + 1, which is the number of unique elements.
*/
