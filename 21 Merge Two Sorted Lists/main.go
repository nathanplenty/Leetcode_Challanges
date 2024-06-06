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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Base cases: if one list is nil, return the other list
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// Recursive cases: merge based on value comparison
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
