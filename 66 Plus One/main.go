package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", plusOne([]int{1, 2, 3}), "\nExpected: [1,2,4]")
	fmt.Println("Output 2:", plusOne([]int{4, 3, 2, 1}), "\nExpected: [4,3,2,2]")
	fmt.Println("Output 3:", plusOne([]int{9}), "\nExpected: [1,0]")

	fmt.Println("STOP")
}

// plusOne adds one to the number represented by the array of digits.
// If the last digit is not 9, it increments it and returns the updated array.
// If the last digit is 9, it handles carry-over by setting it to 0 and propagating the carry.
// If all digits are 9, it adds a new digit '1' at the beginning of the array.
func plusOne(digits []int) []int {
	n := len(digits)

	// Check if the last digit is not 9
	if digits[n-1] != 9 {
		digits[n-1]++ // Increment the last digit
		return digits // Return the updated digits array
	}

	// Handle cases where the last digit is 9
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0 // Set current digit to 0
		} else {
			digits[i]++   // Increment current digit
			return digits // Return the updated digits array
		}
	}

	// If all digits were 9, prepend '1' to the array
	return append([]int{1}, digits...)
}
