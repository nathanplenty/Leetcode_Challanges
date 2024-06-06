package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", isValid("()"), "\nExpected: true")
	fmt.Println("Output 2:", isValid("()[]{}"), "\nExpected: true")
	fmt.Println("Output 3:", isValid("(]"), "\nExpected: false")

	fmt.Println("STOP")
}

// isValid checks if a given string containing brackets is valid.
func isValid(s string) bool {
	// Define a struct to represent a stack of characters
	type Stack struct {
		Chars []rune
	}

	// Define a function to check if a pair of brackets is valid
	isValidPair := func(open, close rune) bool {
		switch close {
		// Check if the closing bracket matches the corresponding opening bracket
		case ')':
			return open == '('
		case ']':
			return open == '['
		case '}':
			return open == '{'
		}
		return false
	}

	// Get the length of the input string
	lengthOfStringIn := len(s)
	// If the string is empty, it's considered valid
	if lengthOfStringIn == 0 {
		return true
	}

	// Initialize an empty stack
	stack := Stack{}

	// Iterate over each character in the string
	for _, char := range s {
		// If the character is an opening bracket, push it onto the stack
		if char == '(' || char == '[' || char == '{' {
			stack.Chars = append(stack.Chars, char)
		} else {
			// If the stack is empty or the closing bracket does not match the last opening bracket, return false
			if len(stack.Chars) == 0 || !isValidPair(stack.Chars[len(stack.Chars)-1], char) {
				return false
			}
			// Remove the matching opening bracket from the stack
			stack.Chars = stack.Chars[:len(stack.Chars)-1]
		}
	}

	// If the stack is empty at the end, all brackets were matched and the string is valid
	return len(stack.Chars) == 0
}
