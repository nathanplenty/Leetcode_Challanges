package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", lengthOfLongestSubstring("abcabcbb"), "\nExpected: 3")
	fmt.Println("Output 2:", lengthOfLongestSubstring("bbbbb"), "\nExpected: 1")
	fmt.Println("Output 3:", lengthOfLongestSubstring("pwwkew"), "\nExpected: 3")

	fmt.Println("STOP")
}

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters
func lengthOfLongestSubstring(s string) int {
	// Print the input string
	fmt.Println(s)

	// Map to store the last positions of each character
	lastSeen := make(map[byte]int)

	// Initialize the length of the longest substring and the starting point of the current substring
	maxLength, start := 0, 0

	// Loop through the string
	for i := 0; i < len(s); i++ {
		// If the character s[i] has been seen and the last seen position is within the current substring
		if pos, found := lastSeen[s[i]]; found && pos >= start {
			// Update the starting point of the current substring
			start = pos + 1
		}
		// Update the last seen position of s[i]
		lastSeen[s[i]] = i
		// Update the length of the longest substring
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	// Return the length of the longest substring without repeating characters
	return maxLength
}
