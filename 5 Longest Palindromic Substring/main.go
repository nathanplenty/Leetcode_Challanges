package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", longestPalindrome("babad"), "\nExpected: bab")
	fmt.Println("Output 2:", longestPalindrome("cbbd"), "\nExpected: bb")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 1 <= s.length <= 1000
	b. s consist of only digits and English letters.
*/

// longestPalindrome O(n^3) finds the longest palindromic substring in a given string.
func longestPalindrome(s string) string {
	var palindromes []string

	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				palindromes = append(palindromes, s[i:j+1])
			}
		}
	}

	var longest string
	for _, p := range palindromes {
		if len(p) > len(longest) {
			longest = p
		}
	}

	return longest
}

/*
Explanation of the function longestPalindrome:

1. Create a slice to store found palindromes:
   - Initialize an empty slice of strings called palindromes.

2. Define a helper function isPalindrome to check if a given string is a palindrome:
   - Iterate through the first half of the string.
   - Compare the character at the current position with the character at the mirrored position from the end.
   - If any characters do not match, return false.
   - If all characters match, return true.

3. Iterate over all possible substrings of s using two nested loops:
   - The outer loop sets the starting index i of the substring.
   - The inner loop sets the ending index j of the substring.
   - Check if the substring s[i:j+1] is a palindrome using the isPalindrome function.
   - If the substring is a palindrome, append it to the palindromes slice.

4. Find the longest palindrome from the palindromes slice:
   - Initialize a variable longest to store the longest palindrome found.
   - Iterate over the palindromes slice.
   - If the current palindrome is longer than the longest, update the longest variable.

5. Return the longest palindrome found.
*/
