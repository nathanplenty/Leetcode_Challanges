package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", isPalindrome(121), "\nExpected: true")
	fmt.Println("Output 2:", isPalindrome(-121), "\nExpected: false")
	fmt.Println("Output 3:", isPalindrome(10), "\nExpected: false")

	fmt.Println("STOP")
}

// Function to check if a given integer is a palindrome.
func isPalindrome(x int) bool {
	// Negative numbers are not palindromes.
	if x < 0 {
		return false
	}

	// Store the original number.
	temp := x
	// Variable to hold the reversed number.
	reversed := 0
	// Loop to reverse the digits of the number.
	for temp != 0 {
		// Get the last digit of the number.
		remainder := temp % 10
		// Append the last digit to the reversed number.
		reversed = reversed*10 + remainder
		// Remove the last digit from the number.
		temp /= 10
	}

	// Check if the original number is equal to the reversed number.
	if x != reversed {
		return false
	}

	// If the number is equal to its reversed version, it is a palindrome.
	return true
}
