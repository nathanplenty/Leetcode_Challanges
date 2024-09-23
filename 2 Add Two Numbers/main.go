package main

import "fmt"

// ListNode given definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("=== START ===")

	// Test 1
	input1 := NewListNode(2, 4, 3)
	input2 := NewListNode(5, 6, 4)
	expected1 := NewListNode(7, 0, 8)
	output1 := addTwoNumbers(input1, input2)
	printResult(input1, input2, output1, expected1)

	// Test 2
	input1 = NewListNode(0)
	input2 = NewListNode(0)
	expected2 := NewListNode(0)
	output2 := addTwoNumbers(input1, input2)
	printResult(input1, input2, output2, expected2)

	// Test 3
	input1 = NewListNode(9, 9, 9, 9, 9, 9, 9)
	input2 = NewListNode(9, 9, 9, 9)
	expected3 := NewListNode(8, 9, 9, 9, 0, 0, 0, 1)
	output3 := addTwoNumbers(input1, input2)
	printResult(input1, input2, output3, expected3)

	fmt.Println("=== STOP ===")
}

// NewListNode creates a linked list from a list of integers.
func NewListNode(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// PrintList prints the elements of a linked list with values separated by commas.
func PrintList(head *ListNode) string {
	if head == nil {
		return "[]"
	}
	result := "["
	result += fmt.Sprint(head.Val)
	head = head.Next
	for head != nil {
		result += ", " + fmt.Sprint(head.Val)
		head = head.Next
	}
	result += "]"
	return result
}

// ANSI-Escape-Codes for colors
const (
	reset = "\033[0m"
	red   = "\033[31m"
	green = "\033[32m"
)

// printResult prints the input, expected, and output in color.
func printResult(input1, input2, output, expected *ListNode) {
	expectedStr := PrintList(expected)
	outputStr := PrintList(output)
	isMatch := compareLists(output, expected)
	fmt.Print("FUNC twoSum()\n- Input 1:  ", PrintList(input1), "\n")
	fmt.Print("- Input 2:  ", PrintList(input2), "\n")
	fmt.Print("- Output:   ")
	if isMatch {
		fmt.Print(green + outputStr + reset + "\n")
	} else {
		fmt.Print(red + outputStr + reset + "\n")
	}
	fmt.Println("- Expected: " + expectedStr)
}

// compareLists compares two linked lists and returns true if they match.
func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// addTwoNumbers O(n) takes two linked lists l1 and l2, each representing a non-negative integer.
// It returns a new linked list representing the sum of l1 and l2.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}
