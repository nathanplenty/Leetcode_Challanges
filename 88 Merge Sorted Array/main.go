package main

import "fmt"

// Information given by Problem.
var nums1 []int
var nums2 []int
var m int
var n int

func main() {
	fmt.Println("START")

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}
	m = 3
	n = 2
	fmt.Println("Output 1:", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("Expected: [1, 2, 2, 3, 5]")

	nums1 = []int{1}
	nums2 = []int{}
	m = 1
	n = 0
	fmt.Println("Output 2:", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("Expected: [1]")

	nums1 = []int{0}
	nums2 = []int{1}
	m = 0
	n = 1
	fmt.Println("Output 3:", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("Expected: [1]")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. nums1.length == m + n
	b. nums2.length == n
	c. 0 <= m, n <= 200
	d. 1 <= m + n <= 200
	e. -10^9 <= nums1[i], nums2[j] <= 10^9
*/

// merge O(m+n) merges two sorted integer arrays nums1 and nums2.
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}

/*
Explanation of the function merge:

1. Initialize three pointers:
   - p1 to the last element of the initial portion of nums1 (m - 1).
   - p2 to the last element of nums2 (n - 1).
   - p to the last element of nums1 (m + n - 1).
2. Iterate while both p1 and p2 are within the valid range:
   - Compare the elements pointed to by p1 and p2.
   - Place the larger element at the current end of nums1 (position p).
   - Move the pointer (p1 or p2) and p one step back accordingly.
3. If there are remaining elements in nums2 after p1 is exhausted:
   - Copy them directly to nums1 starting from position p.
4. No need to handle remaining elements in nums1 as they are already in place.

This approach ensures that all elements are merged in-place efficiently with a time complexity of O(m + n) and a space complexity of O(1).
*/
