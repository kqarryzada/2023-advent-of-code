package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

// isAlmostParallel finds the line in a horizontal matrix that would be parallel
// if one character value was swapped (i.e., a '.' for a '#' or vice versa.)
func isAlmostParallel(index1 int, index2 int, slice []string) int {
	i := index1
	j := index2

	// Check the outer values to ensure symmetry. The unequal strings will hold
	// the values of of mirrored strings that may be mirrored if a character is
	// changed.
	unequalString1 := ""
	unequalString2 := ""
	for {
		if i < 0 || j >= len(slice) {
			break
		}

		if slice[i] != slice[j] {
			if len(unequalString1) == 0 {
				unequalString1 = slice[i]
				unequalString2 = slice[j]
			} else {
				// A second candidate was found, so these indexes cannot refer
				// to a reflection.
				return -1
			}
		}

		i--
		j++
	}

	// When i and j exit the following loop, they will converge to the midway
	// point.
	i = index1 + 1
	j = index2 - 1
	for {
		if i >= j {
			// The midpoint has been reached.
			break
		}

		if slice[i] != slice[j] {
			if len(unequalString1) == 0 {
				unequalString1 = slice[i]
				unequalString2 = slice[j]
			} else {
				// A second candidate was found, so these indexes cannot refer
				// to a reflection.
				return -1
			}
		}

		i++
		j--
	}

	if len(unequalString1) == 0 {
		// An off-by-one reflection was not found.
		return -1
	}

	// Ensure the two unequal strings have a single differing character.
	char := byte('0')
	for k := range unequalString1 {
		if unequalString1[k] != unequalString2[k] {
			if char == '0' {
				char = unequalString1[k]
			} else {
				// These strings have more than one character that are
				// different.
				return -1
			}
		}
	}

	return max(i, j)
}

// calculateValue obtains the numerical value of a matrix given the parallel
// index and whether it corresponds to a row or a column.
func calculateValue(index int, isRow bool) int {
	value := index
	if isRow {
		value *= 100
	}

	return value
}

// findHorizontalParallelLine finds the prallel line of a matrix that is almost
// parallel. If no valid line is found, this function returns -1.
func findHorizontalParallelLine(pattern []string) int {
	for i := range pattern {
		for j := i + 1; j < len(pattern); j += 2 {
			parallelPoint := isAlmostParallel(i, j, pattern)
			if parallelPoint != -1 {
				return parallelPoint
			}
		}
	}

	return -1
}

func invertMatrix(pattern []string) []string {
	inversePattern := make([]string, 0)

	for i := 0; i < len(pattern[0]); i++ {
		inverseString := make([]byte, 0)
		for j := range pattern {
			inverseString = append(inverseString, pattern[j][i])
		}

		inversePattern = append(inversePattern, string(inverseString))
	}

	return inversePattern
}

// computePattern obtains the "value" of a matrix as described by the problem.
// This value is calculated from the number of rows or columns to the left of
// the parallel point (while accounting for a single smudge in the mirror).
func computePattern(pattern []string) int {
	isHorizontal := true
	patternValue := findHorizontalParallelLine(pattern)
	if patternValue == -1 {
		isHorizontal = false
		inversePattern := invertMatrix(pattern)
		patternValue = findHorizontalParallelLine(inversePattern)
		if patternValue == -1 {
			panic("Could not find value for the pattern beginning with: " + pattern[0])
		}
	}

	return calculateValue(patternValue, isHorizontal)
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	sum := 0
	pattern := make([]string, 0)

	for _, line := range fileLines {
		if len(line) != 0 {
			pattern = append(pattern, line)
		} else {
			sum += computePattern(pattern)
			pattern = make([]string, 0)
		}
	}

	if len(pattern) != 0 {
		sum += computePattern(pattern)
	}

	fmt.Printf("The numerical value found from summarizing the notes is %d.\n", sum)
}
