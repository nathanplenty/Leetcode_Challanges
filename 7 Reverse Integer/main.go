package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", reverse(123), "\nExpected: 321")
	fmt.Println("Output 2:", reverse(-123), "\nExpected: -321")
	fmt.Println("Output 3:", reverse(120), "\nExpected: 21")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. -2^31 <= x <= 2^31 - 1
*/

// reverse O(log(x)) converts the integer x into its reverse form.
func reverse(x int) int {
	sx := fmt.Sprintf("%d", x)
	negative := false

	if sx[0] == '-' {
		sx = sx[1:]
		negative = true
	}

	var temp string

	for i := len(sx); i > 0; i-- {
		temp += string(sx[i-1])
	}

	x = 0
	for i := 0; i < len(temp); i++ {
		x = x*10 + int(temp[i]-'0')
	}

	if negative == true {
		x = x * -1
	}

	if x < -(1<<31) || x > (1<<31)-1 {
		return 0
	}

	return x
}

/*
Explanation of the function reverse:

1. Handles negative numbers by first checking if x is negative.
2. Converts x to a string representation and removes the '-' sign if present.
3. Reverses the string representation of x.
4. Converts the reversed string back to an integer.
5. Checks for overflow conditions (if reversed integer exceeds 32-bit signed integer limits).
6. Returns the reversed integer or 0 if overflow occurs.
*/
