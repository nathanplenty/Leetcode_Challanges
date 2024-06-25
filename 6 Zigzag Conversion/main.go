package main

import (
	"fmt"
)

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", convert("PAYPALISHIRING", 3), "\nExpected: PAHNAPLSIIGYIR")
	fmt.Println("Output 2:", convert("PAYPALISHIRING", 4), "\nExpected: PINALSIGYAHRPI")
	fmt.Println("Output 3:", convert("A", 1), "\nExpected: A")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 1 <= s.length <= 1000
	b. s consists of English letters (lower-case and upper-case), ',' and '.'.
	c. 1 <= numRows <= 1000
*/

// convert O(n) converts the input string s into a zigzag pattern with numRows rows.
func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= 1 {
		return s
	}

	cols := len(s)
	if cols > len(s) {
		cols = len(s)
	}

	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	var row, col int
	direction := -1

	for _, char := range s {
		matrix[row][col] = byte(char)

		if row == 0 || row == numRows-1 {
			direction = -direction
		}

		row += direction
		col++
	}

	var result []byte
	for i := 0; i < numRows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != 0 {
				result = append(result, matrix[i][j])
			}
		}
	}

	return string(result)
}

/*
Explanation of the function convert:

1. Convert the input string s into a zigzag pattern with numRows rows.
2. Special cases are handled first:
   - If numRows <= 1 or len(s) <= 1, return s as it is.
3. The function initializes a matrix to hold characters in a Zigzag pattern:
   - numRows rows and the length of the string s columns (or less if numRows is less than the length of s).
4. It fills the matrix by iterating through the string s and placing characters according to the Zigzag pattern:
   - Uses a direction variable to alternate between moving downwards and upwards within the matrix rows.
   - After placing all characters, constructs the result string by appending characters row by row from the matrix to a byte slice.
5. Returns the final result string formed by converting the byte slice to a string.
*/
