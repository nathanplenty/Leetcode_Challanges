package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", climbStairs(2), "\nExpected: 2")
	fmt.Println("Output 2:", climbStairs(3), "\nExpected: 3")

	fmt.Println("STOP")
}

// climbStairs calculates the number of distinct ways to climb a staircase with n steps
// using an iterative approach. The function assumes n is greater than or equal to 1.
func climbStairs(n int) int {
	// If n is less than or equal to 3, return n directly since the number of ways to climb
	// the stairs is equal to n (1 step: 1 way, 2 steps: 2 ways, 3 steps: 3 ways).
	if n <= 3 {
		return n
	}

	// Initialize step1 to the number of ways to climb 2 steps and step2 to the number of ways to climb 3 steps.
	step1, step2 := 2, 3

	// Iterate from 4 up to n, calculating the number of ways to climb each number of steps
	// and updating step1 and step2 accordingly.
	for i := 4; i <= n; i++ {
		// Update step1 to the previous step2 value, and step2 to the sum of the previous step1 and step2 values.
		step1, step2 = step2, step1+step2
	}
	// Return the number of ways to climb n steps.
	return step2
}
