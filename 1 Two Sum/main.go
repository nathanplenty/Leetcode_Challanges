package main

import "fmt"

func main() {
	fmt.Println("=== START ===")

	// Example 1
	input1 := []int{2, 7, 11, 15}
	target1 := 9
	expected1 := []int{0, 1}
	output1 := twoSum(input1, target1)
	printResult(input1, target1, expected1, output1)

	// Example 2
	input2 := []int{3, 2, 4}
	target2 := 6
	expected2 := []int{1, 2}
	output2 := twoSum(input2, target2)
	printResult(input2, target2, expected2, output2)

	// Example 3
	input3 := []int{3, 3}
	target3 := 6
	expected3 := []int{0, 1}
	output3 := twoSum(input3, target3)
	printResult(input3, target3, expected3, output3)

	fmt.Println("=== STOP ===")
}

// ANSI-Escape-Codes for colors
const (
	reset = "\033[0m"
	red   = "\033[31m"
	green = "\033[32m"
)

// formatSlice converts a slice of integers to a comma-separated string within square brackets.
func formatSlice(slice []int) string {
	if len(slice) == 0 {
		return "[]"
	}

	result := "["
	for i, v := range slice {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprint(v)
	}
	result += "]"
	return result
}

// printResult prints the input, target, expected output, and actual output in color.
func printResult(input []int, target int, expected, output []int) {
	isMatch := compareOutputs(output, expected)
	fmt.Print("FUNC twoSum()\n- Input:    ", formatSlice(input), "\n- Target:   ", target, "\n- Expected: ")
	fmt.Print(formatSlice(expected), "\n- Output:   ")
	if isMatch {
		fmt.Print(green + formatSlice(output) + reset + "\n")
	} else {
		fmt.Print(red + formatSlice(output) + reset + "\n")
	}
}

// compareOutputs compares actual and expected outputs and returns true if they match.
func compareOutputs(actual, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

// twoSum finds two numbers in the array that add up to the target value.
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		}
		numMap[num] = i
	}

	return []int{-1, -1}
}
