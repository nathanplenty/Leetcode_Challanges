package main

import "fmt"

func main() {
	fmt.Println("START")

	fmt.Println("Output 1:", intToRoman(3749), "\nExpected: MMMDCCXLIX")
	fmt.Println("Output 2:", intToRoman(58), "\nExpected: LVIII")
	fmt.Println("Output 3:", intToRoman(1994), "\nExpected: MCMXCIV")

	fmt.Println("STOP")
}

/*
Constraints given by the problem:
	a. 1 <= num <= 3999
*/

// intToRoman O(i) converts an integer to a Roman numeral.
func intToRoman(num int) string {
	romanValues := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""

	for _, value := range values {
		for num >= value {
			roman += romanValues[value]
			num -= value
		}
	}

	return roman
}

/*
Explanation of the function intToRoman:

1. Create a map associating integer values with their corresponding Roman numeral symbols, including subtractive combinations.
2. Create a slice of the integer values in descending order.
3. Initialize an empty string `roman` to store the resulting Roman numeral.
4. Iterate over the slice of values:
   a. For each value, while the current integer value (`num`) is greater than or equal to the value, append the corresponding Roman symbol to the `roman` string and subtract the value from `num`.
5. Continue this process until the integer value is reduced to zero.
6. Return the resulting Roman numeral string.
*/
