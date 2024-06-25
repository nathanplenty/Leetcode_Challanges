package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", CODE(), "\nExpected: RETURN")
	fmt.Println("Output 2:", CODE(), "\nExpected: RETURN")
	fmt.Println("Output 3:", CODE(), "\nExpected: RETURN")

	fmt.Println("STOP")
}

// CONSTRAINS

// CODE HEAD COMMENT.
func CODE() string {
	// CODE
	return "RETURN"
}

// CODE EXPLANATION
