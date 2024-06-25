package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", mySqrt(4), "\nExpected: 2")
	fmt.Println("Output 2:", mySqrt(8), "\nExpected: 2")

	fmt.Println("STOP")
}

// mySqrt O(log x) to compute the integer square root of a non-negative integer x.
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 0, x
	var result int

	for left <= right {
		mid := left + (right-left)/2

		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
			result = mid
		} else {
			right = mid - 1
		}
	}

	return result
}

/*
Explanation of the function mySqrt:

1. Handle edge cases where `x` is 0 or 1.
2. Initialize variables `left` and `right` for binary search.
3. Initialize `result` to store the potential result.
4. Perform binary search to find the integer square root:
   a. Calculate the middle point `mid`.
   b. If `mid` is the square root of `x`, return `mid`.
   c. If `mid*mid` is less than `x`, search in the right half and update `result`.
   d. If `mid*mid` is greater than `x`, search in the left half.
5. Return the largest integer whose square is less than or equal to `x`.
*/
