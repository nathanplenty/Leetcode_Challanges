package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", romanToInt("III"), "\nExpected: 3")
	fmt.Println("Output 2:", romanToInt("LVIII"), "\nExpected: 58")
	fmt.Println("Output 3:", romanToInt("MCMXCIV"), "\nExpected: 1994")

	fmt.Println("STOP")
}

// romanToInt (O(i)) input roman symbol number, output value integer
func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	i := 0

	for i < len(s) {
		value := romanValues[s[i]]

		if i+1 < len(s) && romanValues[s[i+1]] > value {
			sum -= value
		} else {
			sum += value
		}

		i++
	}

	if sum == 0 {
		return -1
	}

	return sum
}

/*
Explanation of the function romanToInt:

1. Create a map to associate each Roman symbol with its corresponding integer value.
2. Initialize the sum variable to store the total integer value of the Roman numeral.
3. Initialize the index variable to iterate over the characters of the input string.
4. Iterate over the characters of the input string:
   a. Get the integer value of the current Roman symbol from the map.
   b. Check if the next symbol has a greater value than the current one:
      - If true, subtract the current value from the total sum.
      - If false, add the current value to the total sum.
   c. Move to the next symbol in the input string.
5. Handle the edge case where sum is still 0 after iterating through the string:
   - Return -1 to indicate an error (no valid symbols).
6. Return the total sum of the Roman numeral.
*/
