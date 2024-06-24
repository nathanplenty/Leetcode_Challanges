package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", strStr("sadbutsad", "sad"), "\nExpected: 0")
	fmt.Println("Output 2:", strStr("leetcode", "leeto"), "\nExpected: -1")

	fmt.Println("STOP")
}

// strStr O(n-m+1) finds the first occurrence of 'needle' in 'haystack' and returns its index.
// If 'needle' is not found in 'haystack', it returns -1.
func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		if match {
			return i
		}
	}
	return -1
}

/*
Explanation of the function strStr:

1. Iterate through 'haystack' until the maximum possible start position for 'needle'.
2. For each position 'i' in 'haystack':
   a. Assume 'needle' matches starting from index 'i'.
   b. Check if 'needle' matches starting from index 'i' in 'haystack'.
   c. If any character does not match, set match to false and break the inner loop.
3. If 'needle' matches completely starting from index 'i', return the index 'i'.
4. If 'needle' is not found in 'haystack', return -1.
*/
