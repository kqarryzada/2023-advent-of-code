package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
)

// hasSubstringAtIndex safely checks whether the requested string is present as
// a substring starting at the provided index. For example, for an input of
// ("oneString", 0, "one"), this function will return true.
func hasSubstringAtIndex(line string, i int, number string) bool {
	// Check that there are enough characters remaining in the string.
	if len(line)-i < len(number) {
		return false
	}

	for j := 0; j < len(number); {
		if line[i] != number[j] {
			return false
		}

		i++
		j++
	}

	return true
}

// extractCalibratedValue takes the first and last numerical values in a string
// and returns a two-digit number. For example, "1nineight" returns "18".
func extractCalibratedValue(line string) (calibratedValue int) {
	numberArray := make([]int, 0)

	for i := 0; i < len(line); i++ {
		char := line[i]

		// If the character is a number, add it to the array.
		if char >= '0' && char <= '9' {
			value := int(char) - '0'
			numberArray = append(numberArray, value)
			continue
		}

		// If the index points to a number that's been written out, add it to
		// the array.
		if hasSubstringAtIndex(line, i, "one") {
			numberArray = append(numberArray, 1)
		} else if hasSubstringAtIndex(line, i, "two") {
			numberArray = append(numberArray, 2)
		} else if hasSubstringAtIndex(line, i, "three") {
			numberArray = append(numberArray, 3)
		} else if hasSubstringAtIndex(line, i, "four") {
			numberArray = append(numberArray, 4)
		} else if hasSubstringAtIndex(line, i, "five") {
			numberArray = append(numberArray, 5)
		} else if hasSubstringAtIndex(line, i, "six") {
			numberArray = append(numberArray, 6)
		} else if hasSubstringAtIndex(line, i, "seven") {
			numberArray = append(numberArray, 7)
		} else if hasSubstringAtIndex(line, i, "eight") {
			numberArray = append(numberArray, 8)
		} else if hasSubstringAtIndex(line, i, "nine") {
			numberArray = append(numberArray, 9)
		}
	}

	// Fetch the first and final values in the array.
	var first int
	var last int
	if len(numberArray) == 0 {
		first = 0
		last = 0
	} else if len(numberArray) == 1 {
		first = numberArray[0]
		last = numberArray[0]
	} else {
		first = numberArray[0]
		last = numberArray[len(numberArray)-1]
	}

	// Return the numbers as a two-digit number.
	return first*10 + last
}

func main() {
	sum := 0

	fileLines := fileutils.LoadFile("input.txt")
	for _, line := range fileLines {
		value := extractCalibratedValue(line)
		sum += value
	}

	fmt.Printf("The total sum of the calibrated values is %d.\n", sum)
}
