package main

import (
	"fmt"
	"hash/fnv"
	utils "kqarryzada/advent-of-code-2023/utils"
)

func checksum(input string) uint32 {
	digest := fnv.New32a()
	digest.Write([]byte(input))
	return digest.Sum32()
}

func checksumColumn(columnNumber int, fileLines []string) uint32 {
	bytes := make([]byte, 0)
	for _, line := range fileLines {
		char := line[columnNumber]
		bytes = append(bytes, char)
	}

	return checksum(string(bytes))
}

func parseRows(stringMatrix []string) []uint32 {
	rowHashes := make([]uint32, 0)

	for _, row := range stringMatrix {
		rowHashsum := checksum(row)
		rowHashes = append(rowHashes, rowHashsum)
	}

	return rowHashes
}

func parseColumns(stringMatrix []string) []uint32 {
	columnHashes := make([]uint32, 0)

	for i := range stringMatrix[0] {
		colHashsum := checksumColumn(i, stringMatrix)
		columnHashes = append(columnHashes, colHashsum)
	}

	return columnHashes
}

func isParallel(index1 int, index2 int, slice []uint32) int {
	i := index1
	j := index2

	// Check the outer values to ensure symmetry.
	for {
		if i < 0 || j >= len(slice) {
			break
		}

		if slice[i] != slice[j] {
			return -1
		}

		i--
		j++
	}

	i = index1 + 1
	j = index2 - 1
	for {
		if i >= j {
			// The midpoint has been reached.
			break
		}

		if slice[i] != slice[j] {
			return -1
		}

		i++
		j--
	}

	// When i and j exit the loop, they will converge to the midway point.
	return max(i, j)
}

func calculateValue(index int, isRow bool) int {
	value := index
	if isRow {
		value *= 100
	}

	return value
}

func findParallelLineValue(line []uint32, isRow bool) int {
	for i := range line {
		for j := i + 1; j < len(line); j += 2 {
			if line[i] == line[j] {
				parallelPoint := isParallel(i, j, line)
				if parallelPoint != -1 {
					return calculateValue(parallelPoint, isRow)
				}
			}
		}
	}

	return 0
}

func computePattern(pattern []string) int {
	rowHashsums := parseRows(pattern)
	colHashsums := parseColumns(pattern)
	sum := findParallelLineValue(rowHashsums, true)
	sum += findParallelLineValue(colHashsums, false)

	return sum
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
