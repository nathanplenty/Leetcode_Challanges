package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", addBinary("11", "1"), "\nExpected: 100")
	fmt.Println("Output 2:", addBinary("1010", "1011"), "\nExpected: 10101")

	fmt.Println("STOP")
}

// addBinary O(n) to add two binary strings and return the sum as a binary string.
func addBinary(a string, b string) string {
	var result []byte
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = append([]byte{byte(sum%2) + '0'}, result...)
		carry = sum / 2
	}

	return string(result)
}

/*
Explanation of the function addBinary:

1. Initialize an empty byte slice `result` to store the result binary string.
2. Initialize `carry` to zero.
3. Initialize indices `i` and `j` to point to the last characters of strings `a` and `b`, respectively.
4. Iterate until we have processed all characters of both strings and there's no carry left:
   a. Initialize `sum` with the current value of `carry`.
   b. Add the corresponding bit from string `a` to `sum` if `i` is within bounds.
   c. Add the corresponding bit from string `b` to `sum` if `j` is within bounds.
   d. Append the least significant bit of `sum` (sum % 2) to the beginning of the `result`.
   e. Update `carry` for the next bit addition (sum / 2).
5. Convert the byte slice `result` to a string and return.
*/
