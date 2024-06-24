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

// mergeTwoLists (O(n+m)) merges two sorted linked lists and returns the merged list
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

/*
Explanation of the function mergeTwoLists:

1. If one list is nil, return the other list.
2. Compare the values of the heads of the two lists.
3. Recursively merge the lists based on the value comparison:
   a. If the value of list1's head is less than list2's head, set list1's next to the result of merging the rest of list1 and list2.
   b. Otherwise, set list2's next to the result of merging list1 and the rest of list2.
4. Return the head of the merged list.
*/
