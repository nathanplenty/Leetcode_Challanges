package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", findMedianSortedArrays([]int{1, 3}, []int{2}), "\nExpected: 2.00000")
	fmt.Println("Output 2:", findMedianSortedArrays([]int{1, 2}, []int{3, 4}), "\nExpected: 2.50000")

	fmt.Println("STOP")
}

// findMedianSortedArrays finds the median of the combined sorted arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, k := 0, 0, 0
	merged := make([]int, len(nums1)+len(nums2))

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			i++
		} else {
			merged[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len(nums1) {
		merged[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) {
		merged[k] = nums2[j]
		j++
		k++
	}

	var median float64
	if len(merged)%2 == 0 {
		mid1 := len(merged)/2 - 1
		mid2 := len(merged) / 2
		median = float64(merged[mid1]+merged[mid2]) / 2
	} else {
		mid := len(merged) / 2
		median = float64(merged[mid])
	}

	return median
}

/*
Explanation of the function findMedianSortedArrays:

1. Initialize indices for nums1, nums2, and the merged array.
2. Create a merged array with a length equal to the sum of lengths of nums1 and nums2.
3. Loop to merge nums1 and nums2 until one of them is fully traversed.
   a. If the current element of nums1 is less than the current element of nums2:
      - Add the current element of nums1 to the merged array.
      - Move to the next element in nums1.
   b. Otherwise:
      - Add the current element of nums2 to the merged array.
      - Move to the next element in nums2.
   c. Move to the next position in the merged array.
4. Copy any remaining elements of nums1 into the merged array.
5. Copy any remaining elements of nums2 into the merged array.
6. Declare a variable to hold the median value.
7. Check if the length of the merged array is even:
   a. Calculate the two middle indices.
   b. Calculate the median as the average of the two middle values.
8. Otherwise, if the length of the merged array is odd:
   a. Calculate the middle index.
   b. Calculate the median as the middle value.
9. Return the median value.
*/
