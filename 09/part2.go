package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"slices"
)

// calculateSlope returns an n - 1 length slice containing the differences
// of the input numerical sequence. For example, for an input of
// "1 3 5 7 9", this function will return "2 2 2 2", as well as an isZeroArray
// value of false.
//
// The isZeroArray return value will be true if all the elements in the returned
// slice are zero, indicating that the input sequence was linear.
func calculateSlope(inputSequence *[]int) (slope *[]int, isZeroArray bool) {
	sequence := *inputSequence
	sequenceLength := len(sequence)
	if sequenceLength < 2 {
		panic("Invalid sequence entered.")
	}

	isZero := true
	slopes := make([]int, sequenceLength-1)
	for i := 1; i < sequenceLength; i++ {
		newVal := sequence[i] - sequence[i-1]
		if isZero && newVal != 0 {
			isZero = false
		}
		slopes[i-1] = newVal
	}

	return &slopes, isZero
}

// This function finds the previous value in a sequence of polynomial numbers
// (i.e., the value of the number that would precede the 0th number in the
// sequence).
func calculatePreviousValueInPolynomialSequence(stringSequence string) int {
	sequence := *utils.AsNumericalSlice(stringSequence)

	// Reverse the elements to find the previous value.
	slices.Reverse(sequence)

	// slopeTree holds the lists of slope values. For example:
	// "0 1 4 9"
	// "1 3 5"
	// "2 2"
	slopeTree := make([][]int, 0)
	slopeTree = append(slopeTree, sequence)

	for seq := sequence; ; {
		slopeLine, isZero := calculateSlope(&seq)
		if isZero {
			// A linear sequence has been found.
			break
		}

		slopeTree = append(slopeTree, *slopeLine)
		seq = *slopeLine
	}

	nextValue := slopeTree[len(slopeTree)-1][0]
	for i := len(slopeTree) - 2; i >= 0; i-- {
		seq := slopeTree[i]
		lastValue := seq[len(seq)-1]
		nextValue += lastValue
	}

	return nextValue
}

func main() {
	sum := 0
	fileLines := utils.LoadFile("input.txt")
	for _, line := range fileLines {
		sum += calculatePreviousValueInPolynomialSequence(line)
	}

	fmt.Printf("The sum of all the next values is %d.\n", sum)
}
