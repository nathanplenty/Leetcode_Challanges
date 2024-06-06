package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", longestCommonPrefix([]string{"flower", "flow", "flight"}), "\nExpected: fl")
	fmt.Println("Output 2:", longestCommonPrefix([]string{"dog", "racecar", "car"}), "\nExpected: (nil)")

	fmt.Println("STOP")
}

// longestCommonPrefix takes an array of strings and returns the longest common prefix.
func longestCommonPrefix(strs []string) string {
	// If the input array is empty, return "(nil)" indicating bad input.
	if len(strs) == 0 {
		return "(nil)"
	}

	// Save the first string as the initial prefix.
	prefix := strs[0]

	// Iterate through the remaining strings in the array.
	for i := 1; i < len(strs); i++ {
		// Compare each character of the current string with the corresponding character of the prefix.
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}

		// Update the prefix to contain only the matched characters.
		prefix = prefix[:j]

		// If no characters match, return "(nil)" as there is no common prefix.
		if prefix == "" {
			return "(nil)"
		}
	}

	// Return the longest common prefix found.
	return prefix
}
