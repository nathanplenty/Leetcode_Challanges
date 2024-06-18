package main

import "fmt"

func main() {
	fmt.Println("START")

	// Helper function to print the elements of a linked list
	printList := func(head *ListNode) {
		fmt.Print("[")
		for head != nil {
			fmt.Print(head.Val)
			if head.Next != nil {
				fmt.Print(", ")
			}
			head = head.Next
		}
		fmt.Println("]")
	}

	fmt.Print("Output 1: ")
	printList(deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}))
	fmt.Println("Expected: [1, 2, 3]")

	fmt.Print("Output 2: ")
	printList(deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}))
	fmt.Println("Expected: [1, 2, 3]")

	fmt.Println("STOP")
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to delete duplicates from a sorted singly-linked list.
func deleteDuplicates(head *ListNode) *ListNode {
	// Initialize a pointer to traverse the linked list.
	current := head

	// Traverse the list while there are at least two nodes left to compare.
	for current != nil && current.Next != nil {
		// Check if the current node's value is equal to the next node's value.
		if current.Val == current.Next.Val {
			// If duplicate found, skip the next node by adjusting pointers.
			current.Next = current.Next.Next
		} else {
			// Move to the next node.
			current = current.Next
		}
	}

	return head
}
