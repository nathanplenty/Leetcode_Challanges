package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", lengthOfLastWord("Hello World"), "\nExpected: 5")
	fmt.Println("Output 2:", lengthOfLastWord("   fly me   to   the moon  "), "\nExpected: 4")
	fmt.Println("Output 3:", lengthOfLastWord("luffy is still joyboy"), "\nExpected: 6")

	fmt.Println("STOP")
}

// lengthOfLastWord calculates the length of the last word in the string 's'.
// It returns the length of the last word.
func lengthOfLastWord(s string) int {
	count := 0     // Initialize a counter for the length of the last word
	spacer := true // Flag to track if we are in a sequence of spaces

	// Iterate through the string from the end to the beginning
	for i := len(s); i > 0; i-- {
		// Check if we have encountered a non-space character
		if spacer == false && string(s[i-1]) == " " {
			break // Stop if we encounter a space after the last word
		}
		// Increment the count if the current character is not a space
		if string(s[i-1]) != " " {
			count += 1
			spacer = false // Reset the spacer flag since we found a non-space character
		}
	}

	return count // Return the length of the last word
}
