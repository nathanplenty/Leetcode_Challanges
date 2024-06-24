package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", isValid("()"), "\nExpected: true")
	fmt.Println("Output 2:", isValid("()[]{}"), "\nExpected: true")
	fmt.Println("Output 3:", isValid("(]"), "\nExpected: false")

	fmt.Println("STOP")
}

// isValid (O(n)) checks if a given string containing brackets is valid.
func isValid(s string) bool {
	type Stack struct {
		Chars []rune
	}

	isValidPair := func(open, close rune) bool {
		switch close {
		case ')':
			return open == '('
		case ']':
			return open == '['
		case '}':
			return open == '{'
		}
		return false
	}

	if len(s) == 0 {
		return true
	}

	stack := Stack{}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack.Chars = append(stack.Chars, char)
		} else {
			if len(stack.Chars) == 0 || !isValidPair(stack.Chars[len(stack.Chars)-1], char) {
				return false
			}
			stack.Chars = stack.Chars[:len(stack.Chars)-1]
		}
	}

	return len(stack.Chars) == 0
}

/*
Explanation of the function isValid:

1. Define a struct to represent a stack of characters.
2. Define a function to check if a pair of brackets is valid.
3. If the input string is empty, it's considered valid.
4. Initialize an empty stack.
5. Iterate over each character in the string:
   a. If the character is an opening bracket, push it onto the stack.
   b. If the stack is empty or the closing bracket does not match the last opening bracket, return false.
   c. Remove the matching opening bracket from the stack.
6. If the stack is empty at the end, all brackets were matched and the string is valid.
*/
