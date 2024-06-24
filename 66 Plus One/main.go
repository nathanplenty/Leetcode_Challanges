package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", plusOne([]int{1, 2, 3}), "\nExpected: [1,2,4]")
	fmt.Println("Output 2:", plusOne([]int{4, 3, 2, 1}), "\nExpected: [4,3,2,2]")
	fmt.Println("Output 3:", plusOne([]int{9}), "\nExpected: [1,0]")

	fmt.Println("STOP")
}

// plusOne O(n) adds one to the number represented by the array of digits.
// If the last digit is not 9, it increments it and returns the updated array.
// If the last digit is 9, it handles carry-over by setting it to 0 and propagating the carry.
// If all digits are 9, it adds a new digit '1' at the beginning of the array.
func plusOne(digits []int) []int {
	n := len(digits)
	if digits[n-1] != 9 {
		digits[n-1]++
		return digits
	}
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return append([]int{1}, digits...)
}

/*
Explanation of the function plusOne:

1. Get the length of the digits array `n`.
2. If the last digit is not 9, increment it and return the updated array.
3. If the last digit is 9, iterate through the array from the end to the beginning:
   a. If the current digit is 9, set it to 0.
   b. If the current digit is not 9, increment it and return the updated array.
4. If all digits are 9, prepend '1' to the array and return it.
*/
