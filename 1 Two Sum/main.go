package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", twoSumQuad([]int{2, 7, 11, 15}, 9), "\nExpected: [0 1]")
	fmt.Println("Output 2:", twoSumLin([]int{3, 2, 4}, 6), "\nExpected: [1 2]")
	fmt.Println("Output 3:", twoSumQuad([]int{3, 3}, 6), "\nExpected: [0 1]")

	fmt.Println("STOP")
}

// twoSumQuad finds two numbers in the array that add up to the target value using a quadratic time complexity approach (O(n^2)).
func twoSumQuad(nums []int, target int) []int {
	fmt.Println("FUNC twoSumQuad()")
	// Iterate through each element in the array.
	for i := 0; i < len(nums); i++ {
		// Calculate the complement of the current element.
		complement := target - nums[i]
		// Check the subsequent elements to find the complement.
		for j := i + 1; j < len(nums); j++ {
			// If the complement is found, return the indices of the two elements.
			if nums[j] == complement {
				return []int{i, j}
			}
		}
	}
	// If no pair is found, return [-1, -1].
	return []int{-1, -1}
}

// twoSumLin finds two numbers in the array that add up to the target value using a linear time complexity approach (O(n)).
func twoSumLin(nums []int, target int) []int {
	fmt.Println("FUNC twoSumLin()")
	// Create a hashmap to store the value and its index.
	numMap := make(map[int]int)

	// Iterate through each element in the array.
	for i, num := range nums {
		// Calculate the complement of the current element.
		complement := target - num
		// Check if the complement is already in the hashmap.
		if index, found := numMap[complement]; found {
			// If the complement is found, return the indices of the two elements.
			return []int{index, i}
		}
		// Add the current element and its index to the hashmap.
		numMap[num] = i
	}

	// If no pair is found, return [-1, -1].
	return []int{-1, -1}
}
