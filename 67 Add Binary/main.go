package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", addBinary("11", "1"), "\nExpected: 100")
	fmt.Println("Output 2:", addBinary("1010", "1011"), "\nExpected: 10101")

	fmt.Println("STOP")
}

// Function to add two binary strings and return the sum as a binary string.
func addBinary(a string, b string) string {
	// Initialize an empty byte slice to store the result binary string.
	var result []byte

	// Initialize carry to zero.
	carry := 0

	// Initialize indices i and j to point to the last characters of strings a and b, respectively.
	i, j := len(a)-1, len(b)-1

	// Iterate until we have processed all characters of both strings and there's no carry left.
	for i >= 0 || j >= 0 || carry > 0 {
		// Initialize sum with the current value of carry.
		sum := carry

		// Add the corresponding bit from string a to sum if i is within bounds.
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		// Add the corresponding bit from string b to sum if j is within bounds.
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// Append the least significant bit of sum (sum % 2) to the beginning of the result.
		result = append([]byte{byte(sum%2) + '0'}, result...)

		// Update carry for the next bit addition (sum / 2).
		carry = sum / 2
	}

	// Convert the byte slice result to a string and return.
	return string(result)
}
