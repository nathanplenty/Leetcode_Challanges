package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", lengthOfLastWord("Hello World"), "\nExpected: 5")
	fmt.Println("Output 2:", lengthOfLastWord("   fly me   to   the moon  "), "\nExpected: 4")
	fmt.Println("Output 3:", lengthOfLastWord("luffy is still joyboy"), "\nExpected: 6")

	fmt.Println("STOP")
}

// lengthOfLastWord O(n) calculates the length of the last word in the string 's'.
// It returns the length of the last word.
func lengthOfLastWord(s string) int {
	count := 0
	spacer := true

	for i := len(s); i > 0; i-- {
		if spacer == false && s[i-1] == ' ' {
			break
		}
		if s[i-1] != ' ' {
			count++
			spacer = false
		}
	}

	return count
}

/*
Explanation of the function lengthOfLastWord:

1. Initialize a counter `count` for the length of the last word.
2. Initialize a flag `spacer` to track if we are in a sequence of spaces.
3. Iterate through the string `s` from the end to the beginning:
   a. If `spacer` is false and the current character is a space, break the loop.
   b. If the current character is not a space, increment the counter `count` and set `spacer` to false.
4. Return the value of `count`, which represents the length of the last word.
*/
