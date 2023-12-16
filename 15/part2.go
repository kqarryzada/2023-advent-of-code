package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"strconv"
	"strings"
)

type operation struct {
	label       string
	opType      operationType
	focalLength int
}

type operationType int

const (
	REPLACE operationType = iota
	REMOVE
)

func sequenceHash(sequence string) int {
	hash := 0
	for _, char := range sequence {
		hash += int(char)
		hash *= 17
		hash %= 256
	}

	return hash
}

// parseOperations converts a list of raw string operation values into a slice
// of operation objects.
func parseOperations(stringOperations []string) []operation {
	operations := make([]operation, 0)
	for _, stringOp := range stringOperations {
		op := new(operation)

		if strings.Contains(stringOp, "=") {
			values := strings.Split(stringOp, "=")
			op.label = values[0]
			op.opType = REPLACE
			lensValue, err := strconv.Atoi(values[1])
			if err != nil {
				panic("Error when parsing value: " + stringOp)
			}
			op.focalLength = lensValue
		} else if strings.HasSuffix(stringOp, "-") {
			op.label = strings.Split(stringOp, "-")[0]
			op.opType = REMOVE
			op.focalLength = -1
		} else {
			panic("Unexpected operation format: " + stringOp)
		}

		operations = append(operations, *op)
	}

	return operations
}

// removeLensFromBox removes an operation from a slice. If the requested
// operation label is not present in the slice, the original slice is returned.
func removeLensFromBox(label string, box []operation) []operation {
	requestedLensIndex := -1
	for i, lens := range box {
		if lens.label == label {
			requestedLensIndex = i
			break
		}
	}

	if requestedLensIndex == -1 {
		return box
	}

	if requestedLensIndex == 0 {
		return box[1:]
	}

	if requestedLensIndex == len(box)-1 {
		return box[:len(box)-1]
	}

	return append(box[:requestedLensIndex], box[requestedLensIndex+1:]...)
}

// replaceLensInBox replaces an labelled operation within a slice. If the
// requested operation label is not present, it will be appended to the end of
// the lice.
func replaceLensInBox(op operation, box []operation) []operation {
	requestedLensIndex := -1
	for i, lens := range box {
		if lens.label == op.label {
			requestedLensIndex = i
			break
		}
	}

	if requestedLensIndex == -1 {
		return append(box, op)
	}

	box[requestedLensIndex] = op
	return box
}

// process consumes a list of operations and arranges them as a list of boxes,
// where each box contains a list of lenses.
func process(operations []operation) [][]operation {
	// The boxList stores the arrangement of lenses within the 256 boxes. Since
	// the data we need in the boxes is already contained in the operations,
	// each box holds a slice of operations.
	boxList := make([][]operation, 256)

	for _, op := range operations {
		boxID := sequenceHash(op.label)
		box := boxList[boxID]

		var newLensOrder []operation
		if op.opType == REMOVE {
			newLensOrder = removeLensFromBox(op.label, box)
		} else {
			newLensOrder = replaceLensInBox(op, box)
		}
		boxList[boxID] = newLensOrder
	}

	return boxList
}

func run(operations []operation) int {
	boxList := process(operations)

	sum := 0
	for i, lensList := range boxList {
		for j, lens := range lensList {
			boxNumber := i + 1
			lensNumber := j + 1
			focusingPower := boxNumber * lensNumber * lens.focalLength
			sum += focusingPower
		}
	}

	return sum
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	if len(fileLines) != 1 {
		panic("Unexpected input file format.")
	}

	stringSequence := strings.Split(fileLines[0], ",")
	initializationSequence := parseOperations(stringSequence)
	sum := run(initializationSequence)

	fmt.Printf("The total focusing power is %d.\n", sum)
}
