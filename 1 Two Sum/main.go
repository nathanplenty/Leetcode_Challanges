package main

/*
Two Sum:
	Given an array of integers nums and an integer target,
	return indices of the two numbers such that they add up to target.
	You may assume that each input would have exactly one solution,
	and you may not use the same element twice.
	You can return the answer in any order.
*/

import "fmt"

func main() {
	fmt.Println("=== START ===")

	fmt.Println("FUNC twoSumLin()\n- Output 1:", twoSum([]int{2, 7, 11, 15}, 9), "\n- Expected: [0 1]")
	fmt.Println("FUNC twoSumLin()\n- Output 2:", twoSum([]int{3, 2, 4}, 6), "\n- Expected: [1 2]")
	fmt.Println("FUNC twoSumLin()\n- Output 3:", twoSum([]int{3, 3}, 6), "\n- Expected: [0 1]")

	fmt.Println("=== STOP ===")
}

/*
Constraints given by the problem:
	a. 2 <= nums.length <= 10^4
	b. -10^9 <= nums[i] <= 10^9
	c. -10^9 <= target <= 10^9
	d. Only one valid answer exists
*/

// twoSum O(n) finds two numbers in the array that add up to the target value.
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return []int{-1, -1}
}

/*
Explanation of the function twoSumLin:

1. Create a hashmap to store the value and its index.
2. Iterate through each element in the array.
   a. Calculate the complement of the current element.
   b. Check if the complement is already in the hashmap.
   c. If the complement is found, return the indices of the two elements.
   d. Add the current element and its index to the hashmap.
3. If no pair is found, return [-1, -1].
*/
