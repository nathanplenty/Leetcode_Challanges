package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", lengthOfLongestSubstring("abcabcbb"), "\nExpected: 3")
	fmt.Println("Output 2:", lengthOfLongestSubstring("bbbbb"), "\nExpected: 1")
	fmt.Println("Output 3:", lengthOfLongestSubstring("pwwkew"), "\nExpected: 3")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 0 <= s.length <= 5 * 10^4
	b. s consists of English letters, digits, symbols and spaces.
*/

// lengthOfLongestSubstring O(n) finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLength, start := 0, 0

	for i := 0; i < len(s); i++ {
		if pos, found := lastSeen[s[i]]; found && pos >= start {
			start = pos + 1
		}
		lastSeen[s[i]] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

/*
Explanation of the function lengthOfLongestSubstring:

1. Map Initialization:
   - Create a map `lastSeen` to store the last positions of each character in the string `s`.
2. Variables Initialization:
   - Initialize `maxLength` to store the length of the longest substring without repeating characters.
   - Initialize `start` to keep track of the starting index of the current substring.
3. Loop through the string:
   - Iterate through each character `s[i]` in the string.
4. Handling Duplicate Characters:
   - Check if the character `s[i]` has been seen before and if its last seen position (`pos`) is within the current substring (`start` to `i`).
   - If yes, update the starting index of the current substring (`start`) to `pos + 1`.
5. Update Last Seen Position:
   - Update the last seen position of the character `s[i]` to `i` in the `lastSeen` map.
6. Update Maximum Length:
   - Calculate the length of the current substring (`i - start + 1`).
   - Update `maxLength` if the current substring length is greater than `maxLength`.
7. Return Result:
   - Return `maxLength`, which represents the length of the longest substring without repeating characters.
*/
