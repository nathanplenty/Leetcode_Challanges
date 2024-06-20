package main

import "fmt"

// ListNode specified definition for singly-linked list.
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
	printList(addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}))
	fmt.Println("Expected: [7, 0, 8]")

	fmt.Print("Output 2: ")
	printList(addTwoNumbers(&ListNode{0, nil}, &ListNode{0, nil}))
	fmt.Println("Expected: [0]")

	fmt.Print("Output 3: ")
	printList(addTwoNumbers(&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}))
	fmt.Println("Expected: [8, 9, 9, 9, 0, 0, 0, 1]")

	fmt.Println("STOP")
}

// addTwoNumbers takes two linked lists l1 and l2, each representing a non-negative integer.
// It returns a new linked list representing the sum of l1 and l2 (O(n)).
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

/*
Explanation of the function addTwoNumbers:

1. Dummy Node Initialization:
   - Create a dummy node to facilitate the creation of the result list.
   - Initialize a current pointer to keep track of the current node in the result list.
   - Initialize carry to 0, which initially represents no carry over.
2. Iteration Through Lists:
   - Loop continues as long as there are nodes in either l1, l2, or there's a carry.
   - Compute sum as carry (initially 0).
3. Adding Values:
   - If l1 has a node, add its value to sum and move to the next node.
   - If l2 has a node, add its value to sum and move to the next node.
4. Carry Calculation:
   - Update carry to sum / 10 for the next iteration.
5. Result List Construction:
   - Create a new node with value sum % 10 and link it to the current node's Next.
   - Move the current pointer to the newly added node.
6. Return Result:
   - Return the head of the resulting list, which is dummy.Next.
*/
