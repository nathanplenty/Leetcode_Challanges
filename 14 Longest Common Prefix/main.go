package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", longestCommonPrefix([]string{"flower", "flow", "flight"}), "\nExpected: fl")
	fmt.Println("Output 2:", longestCommonPrefix([]string{"dog", "racecar", "car"}), "\nExpected: (nil)")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 1 <= strs.length <= 200
	b. 0 <= strs[i].length <= 200
	c. strs[i] consists of only lowercase English letters.
*/

// longestCommonPrefix O(n*m) takes an array of strings and returns the longest common prefix.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "(nil)"
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}
		prefix = prefix[:j]

		if prefix == "" {
			return "(nil)"
		}
	}

	return prefix
}

/*
Explanation of the function longestCommonPrefix:

1. If the input array is empty, return "(nil)" indicating bad input.
2. Save the first string as the initial prefix.
3. Iterate through the remaining strings in the array:
   a. Compare each character of the current string with the corresponding character of the prefix.
   b. Update the prefix to contain only the matched characters.
   c. If no characters match, return "(nil)" as there is no common prefix.
4. Return the longest common prefix found.
*/
