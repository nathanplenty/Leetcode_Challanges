package main

import "fmt"

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
	fmt.Println("Output 1: ", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("\nExpected: [1, 2, 2, 3, 5, 6]")

	nums1 = []int{1}
	nums2 = []int{}
	m = 1
	n = 0
	fmt.Println("Output 2: ", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("\nExpected: [1]")

	nums1 = []int{0}
	nums2 = []int{1}
	m = 0
	n = 1
	fmt.Println("Output 3: ", nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	fmt.Println("\nExpected: [1]")

	fmt.Println("STOP")
}

// CODE HEAD COMMENT.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// don't write code, will do later
	fmt.Println(nums1, m, nums2, n)
}
