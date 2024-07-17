package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", threeSumClosest([]int{-1, 2, 1, -4}, 1), "\nExpected: 2")
	fmt.Println("Output 2:", threeSumClosest([]int{0, 0, 0}, 1), "\nExpected: 0")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 3 <= nums.length <= 500
	b. -1000 <= nums[i] <= 1000
	c. -10^4 <= target <= 10^4
*/

// threeSumClosest O(n^2) threeSumClosest finds the sum of three integers in nums that is closest to the target value.
func threeSumClosest(nums []int, target int) int {
	n := len(nums)

	if n < 3 {
		return 0
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

	closest := 1<<31 - 1

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := sum - target
			if diff < 0 {
				diff = -diff
			}
			if diff < abs(closest-target) {
				closest = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closest
}

/*
Explanation of the function threeSumClosest:

1. The function starts by checking if the length of the input array `nums` is less than 3. If so, it returns `0` because there aren't enough elements to form a triplet.
2. An insertion sort function is defined and used to sort the array. Sorting the array helps to efficiently use the two-pointer technique to find the closest sum to the target.
3. After sorting, the function initializes a variable `closest` to a very large number (`1<<31 - 1`), which will keep track of the closest sum found.
4. The function then iterates through each element in the sorted array using an index `i`. For each element, it uses two pointers, `left` and `right`, which start just after the current element and at the end of the array, respectively.
5. Within this loop, it calculates the sum of the elements at the indices `i`, `left`, and `right`. It computes the absolute difference between this sum and the `target`.
6. If the absolute difference between the current sum and the target is smaller than the difference between the `closest` sum and the target, the `closest` value is updated to the current sum.
7. The pointers are adjusted based on the current sum:
   - If the sum is less than the target, the `left` pointer is moved to the right to increase the sum.
   - If the sum is greater than the target, the `right` pointer is moved to the left to decrease the sum.
   - If the sum equals the target, the function returns this sum immediately as it is the closest possible.
8. After all iterations, the function returns the `closest` sum found, which is the sum of three integers closest to the target value.
*/
