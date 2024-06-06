package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", romanToInt("III"))
	fmt.Println("Output 2:", romanToInt("LVIII"))
	fmt.Println("Output 3:", romanToInt("MCMXCIV"))

	fmt.Println("STOP")
}

// romanToInt input roman symbol number, output value integer
func romanToInt(s string) int {
	// Create a map to associate each Roman symbol with its corresponding integer value.
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Initialize the sum variable to store the total integer value of the Roman numeral.
	sum := 0

	// Initialize the index variable to iterate over the characters of the input string.
	i := 0

	// Iterate over the characters of the input string.
	for i < len(s) {
		// Get the integer value of the current Roman symbol from the map.
		value := romanValues[s[i]]

		// If the next symbol has a greater value than the current one, subtract the current value from the total.
		// Otherwise, add the current value to the total.
		if i+1 < len(s) && romanValues[s[i+1]] > value {
			sum -= value
		} else {
			sum += value
		}

		// Move to the next symbol in the input string.
		i++
	}

	// If the sum is still 0, it means there were no symbols in the input string.
	// Return -1 in this case to indicate an error.
	if sum == 0 {
		return -1
	}

	// Return the total sum of the Roman numeral.
	return sum
}
