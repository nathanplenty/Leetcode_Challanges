package main

import "fmt"

// ListNode given definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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

// deleteDuplicates O(n) to delete duplicates from a sorted singly-linked list.
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

/*
Explanation of the function deleteDuplicates:

1. Initialize a pointer to traverse the linked list.
2. Traverse the list while there are at least two nodes left to compare:
   a. Check if the current node's value is equal to the next node's value.
   b. If a duplicate is found, skip the next node by adjusting pointers.
   c. If no duplicate is found, move to the next node.
3. Return the head of the modified linked list.
*/
