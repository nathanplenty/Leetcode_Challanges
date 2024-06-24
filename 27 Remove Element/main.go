package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", removeElement([]int{3, 2, 2, 3}, 3), "\nExpected: RETURN")
	fmt.Println("Output 2:", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), "\nExpected: RETURN")

	fmt.Println("STOP")
}

// removeElement O(n) removes all occurrences of 'val' from the slice 'nums' in place.
// It returns the length of the modified slice after removal.
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/*
Explanation of the function removeElement:

1. Initialize the slow pointer at index 0.
2. Iterate with the fast pointer from index 0 to the end of the array:
   a. If `nums[fast]` is not equal to `val`, copy `nums[fast]` to `nums[slow]` and increment `slow`.
3. Return `slow`, which is the length of the modified slice `nums` without the elements equal to `val`.
*/
