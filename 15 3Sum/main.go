package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", threeSum([]int{-1, 0, 1, 2, -1, -4}), "\nExpected: [[-1,-1,2],[-1,0,1]]")
	fmt.Println("Output 2:", threeSum([]int{0, 0, 1}), "\nExpected: []")
	fmt.Println("Output 3:", threeSum([]int{0, 0, 0}), "\nExpected: [[0,0,0]]")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 3 <= nums.length <= 3000
	b. -10^5 <= nums[i] <= 10^5
*/

// threeSum O(n^2) function finds all unique triplets in the array which gives the sum of zero.
func threeSum(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 3 {
		return result
	}

	insertionSort := func(nums []int) {
		for i := 1; i < len(nums); i++ {
			key := nums[i]
			j := i - 1
			for j >= 0 && nums[j] > key {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = key
		}
	}

	insertionSort(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

/*
Explanation of the function threeSum:

1. The function starts by checking if the length of the input array is less than 3. If so, it returns an empty result as there cannot be any triplets.
2. An insertion sort function is defined and used to sort the array. Sorting helps to efficiently find and avoid duplicate triplets.
3. After sorting the array, we iterate through it. For each element, we use two pointers, `left` and `right`, to find pairs that sum up to the negative value of the current element.
4. If a valid triplet is found, it is added to the result. To avoid duplicates, we skip over duplicate elements for both the left and right pointers.
5. If the sum of the triplet is less than zero, we move the left pointer to the right to increase the sum. If the sum is greater than zero, we move the right pointer to the left to decrease the sum.
*/
