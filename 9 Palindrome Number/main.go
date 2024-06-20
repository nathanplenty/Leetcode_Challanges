package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", isPalindrome(121), "\nExpected: true")
	fmt.Println("Output 2:", isPalindrome(-121), "\nExpected: false")
	fmt.Println("Output 3:", isPalindrome(10), "\nExpected: false")

	fmt.Println("STOP")
}

// isPalindrome (O(d)) checks if a given integer is a palindrome.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var digits []int

	temp := x
	for temp > 0 {
		digit := temp % 10
		digits = append(digits, digit)
		temp /= 10
	}

	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}

/*
Explanation of the function isPalindrome:

1. Check if the number x is lower than 0:
   - If x is negative, return false because negative numbers are not palindromes.
2. Initialize an empty slice called digits to store the digits of x in reverse order.
3. Extract digits from x and store them in reverse order in the digits slice:
   - Use a loop that continues until x is greater than 0.
   - Calculate the last digit of x using x % 10 and append it to the digits slice.
   - Remove the last digit from x by dividing x by 10.
4. Check for palindrome property by comparing digits from both ends towards the center:
   - Loop through the first half of the digits slice.
   - Compare digits[i] with digits[len(digits)-1-i].
   - If any pair of digits doesn't match, return false.
5. If all comparisons are equal, return true because x is a palindrome.
*/
