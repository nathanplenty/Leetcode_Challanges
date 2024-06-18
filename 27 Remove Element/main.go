package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", removeElement([]int{3, 2, 2, 3}, 3), "\nExpected: RETURN")
	fmt.Println("Output 2:", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), "\nExpected: RETURN")

	fmt.Println("STOP")
}

// removeElement removes all occurrences of 'val' from the slice 'nums' in place.
// It returns the length of the modified slice after removal.
func removeElement(nums []int, val int) int {
	// Iterate through the slice 'nums'
	for i := 0; i < len(nums); {
		// If current element 'nums[i]' is equal to 'val', remove it from 'nums'
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:]...) // Remove element at index 'i' from 'nums'
		} else {
			i++ // Increment 'i' only when 'nums[i]' is not equal to 'val'
		}
	}

	// Return the length of the modified slice 'nums'
	return len(nums)
}
