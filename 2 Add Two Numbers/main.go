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

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers takes two linked lists l1 and l2, each representing a non-negative integer.
// It returns a new linked list representing the sum of l1 and l2.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy node to facilitate creation of result list
	current := dummy     // Pointer to current node in the result list
	carry := 0           // Carry over value initially zero

	// Iterate through lists l1 and l2 and continue if there's a carry
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // Initialize sum with carry value

		// Add values from l1 and move to the next node
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Add values from l2 and move to the next node
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10                        // Update carry for the next iteration
		current.Next = &ListNode{Val: sum % 10} // Create new node with sum % 10 and link it to the result list
		current = current.Next                  // Move current pointer to the next node in the result list
	}

	return dummy.Next // Return the next node of dummy, which is the head of the resulting list
}
