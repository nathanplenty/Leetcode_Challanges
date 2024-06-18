package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", strStr("sadbutsad", "sad"), "\nExpected: 0")
	fmt.Println("Output 2:", strStr("leetcode", "leeto"), "\nExpected: -1")

	fmt.Println("STOP")
}

// strStr finds the first occurrence of 'needle' in 'haystack' and returns its index.
// If 'needle' is not found in 'haystack', it returns -1.
func strStr(haystack string, needle string) int {
	// Iterate through 'haystack' until the maximum possible start position for 'needle'
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true

		// Check if 'needle' matches starting from index 'i' in 'haystack'
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}

		// If 'needle' matches completely starting from index 'i', return the index 'i'
		if match {
			return i
		}
	}

	// If 'needle' is not found in 'haystack', return -1
	return -1
}
