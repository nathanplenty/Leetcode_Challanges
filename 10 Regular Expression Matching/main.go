package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", isMatch("aa", "a"), "\nExpected: false")
	fmt.Println("Output 2:", isMatch("aa", "a*"), "\nExpected: true")
	fmt.Println("Output 3:", isMatch("ab", ".*"), "\nExpected: true")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 1 <= s.length <= 20
	b. 1 <= p.length <= 20
	c. s contains only lowercase English letters.
	d. p contains only lowercase English letters, '.', and '*'.
	e. It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
*/

// isMatch O(2^(m+n)) checks if the given string s matches the pattern p.
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

/*
Explanation of the function isMatch:

1. Base case: If the pattern p is empty, return true if the string s is also empty; otherwise, return false.
2. Check if the first character of s matches the first character of p or if p[0] is a '.' which matches any character:
   - Set firstMatch to true if there is a match.
3. Handle the '*' wildcard in the pattern:
   - If the second character of p is '*', there are two possibilities:
     a. Skip the '*' and the preceding element (p[2:]) and call isMatch with the rest of the pattern.
     b. Use the '*' to match the first character of s and continue matching the rest of s with the same pattern.
4. If there is no '*', recursively check the next characters of s and p:
   - Call isMatch with the next characters of s (s[1:]) and p (p[1:]) if firstMatch is true.
*/
