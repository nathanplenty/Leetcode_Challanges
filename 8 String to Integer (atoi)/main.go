package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", myAtoi("42"), "\nExpected: 42")
	fmt.Println("Output 2:", myAtoi(" -042"), "\nExpected: -42")
	fmt.Println("Output 3:", myAtoi("1337c0d3"), "\nExpected: 1337")
	fmt.Println("Output 4:", myAtoi("0-1"), "\nExpected: 0")
	fmt.Println("Output 5:", myAtoi("words and 987"), "\nExpected: 0")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 0 <= s.length <= 200
	b. s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/

// myAtoi O(n) converts the string s into an integer.
func myAtoi(s string) int {
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	var negative bool
	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		negative = false
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	var x int
	for len(s) > 0 && (s[0] >= '0' && s[0] <= '9') {
		fmt.Println(x)
		x = (x * 10) + int(s[0]-'0')
		s = s[1:]
		if x >= (1 << 31) {
			x = (1 << 31) - 1
			if negative == true {
				x += 1
			}
			break
		}
	}

	if negative == true {
		x = -x
	}

	return x
}

/*
Explanation of the function myAtoi:

1. Trims leading whitespace characters from the input string s.
2. Checks for optional leading '+' or '-' signs to determine the sign of the integer.
3. Parses numeric digits until a non-digit character is encountered or until the integer overflows.
4. Handles overflow conditions by returning the maximum or minimum 32-bit signed integer.
5. Returns the parsed integer with the correct sign based on the input.
*/
