package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", longestPalindrome("babad"), "\nExpected: bab")
	fmt.Println("Output 2:", longestPalindrome("cbbd"), "\nExpected: bb")

	fmt.Println("STOP")
}

// longestPalindrome O(n^2) finds the longest palindromic substring in a given string.
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var longest string

	isPalindrome := func(s string) bool {
		runes := []rune(s)
		length := len(runes)

		for i := 0; i < length/2; i++ {
			if runes[i] != runes[length-1-i] {
				return false
			}
		}

		return true
	}

	expandAroundCenter := func(left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return s[left+1 : right]
	}

	for i := 0; i < len(s); i++ {
		odd := expandAroundCenter(i, i)
		even := expandAroundCenter(i, i+1)

		if isPalindrome(odd) && len(odd) > len(longest) {
			longest = odd
		}
		if isPalindrome(even) && len(even) > len(longest) {
			longest = even
		}
	}

	return longest
}

/*
Explanation of the function longestPalindrome:

1. Check if the input string length is less than or equal to 1:
   - If true, return the input string as it is already a palindrome.

2. Initialize a variable longest to store the longest palindrome found.

3. Define a helper function isPalindrome to check if a given string is a palindrome:
   - Convert the string to a slice of runes to properly handle Unicode characters.
   - Iterate through the first half of the runes.
   - Compare the rune at the current position with the rune at the mirrored position from the end.
   - If any runes do not match, return false.
   - If all runes match, return true.

4. Define a helper function expandAroundCenter to find the longest palindrome centered at given indices:
   - Expand outwards from the center indices as long as the characters at the left and right indices are equal.
   - Return the substring between the final left and right indices.

5. Iterate over each character in the string to check for palindromes centered at each character (odd length) and between each pair of characters (even length):
   - Call expandAroundCenter with the current index for odd length palindromes.
   - Call expandAroundCenter with the current and next index for even length palindromes.
   - If the palindrome found is longer than the current longest palindrome, update longest.

6. Return the longest palindrome found.
*/
