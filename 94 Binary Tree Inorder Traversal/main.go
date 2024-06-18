package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", CODE(), "\nExpected: RETURN")
	fmt.Println("Output 2:", CODE(), "\nExpected: RETURN")
	fmt.Println("Output 3:", CODE(), "\nExpected: RETURN")

	fmt.Println("STOP")
}

// CODE HEAD COMMENT.
func CODE() string {
	//
	return "RETURN"
}
