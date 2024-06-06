package main

import "fmt"

// ListNode struct represents a node in a linked list
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
	printList(mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
	fmt.Println("Expected: [1, 1, 2, 3, 4, 4]")

	fmt.Print("Output 2: ")
	printList(mergeTwoLists(nil, nil))
	fmt.Println("Expected: []")

	fmt.Print("Output 3: ")
	printList(mergeTwoLists(nil, &ListNode{0, nil}))
	fmt.Println("Expected: [0]")

	fmt.Println("STOP")
}

// mergeTwoLists merges two sorted linked lists and returns the merged list
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Base cases: if one list is nil, return the other list
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// Recursive cases: merge based on value comparison
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
