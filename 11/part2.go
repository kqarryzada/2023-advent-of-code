package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

// This value corresponds to the number of rows/columns that should be inserted
// in the place of an empty row or column.
var DISTANCE_MULTIPLIER = 1000000

type coordinate struct {
	row int
	col int
}

// This function parses the input file to extract a list of coordinates that
// correspond to the locations of galaxies.
func parse(fileLines []string) []*coordinate {
	coordinateList := make([]*coordinate, 0)

	row := 0
	for _, line := range fileLines {
		blankLine := true
		for col, char := range line {
			if char == '#' {
				blankLine = false
				galaxy := &coordinate{row, col}
				coordinateList = append(coordinateList, galaxy)
			}

		}

		if blankLine {
			// There were no galaxies in this row.
			row += DISTANCE_MULTIPLIER - 1
		}
		row++
	}

	// Assemble a list of blank columns.
	blankColumnList := make([]int, 0)
	for j := range fileLines[0] {
		blankColumn := true
		for i := range fileLines {
			if fileLines[i][j] != '.' {
				blankColumn = false
				break
			}
		}
		if blankColumn {
			blankColumnList = append(blankColumnList, j)
		}
	}

	// Iterate through the assembled coordinates and update the column values.
	for _, coord := range coordinateList {
		originalColumn := coord.col
		for _, blankColumn := range blankColumnList {
			if originalColumn >= blankColumn {
				coord.col += DISTANCE_MULTIPLIER - 1
			}
		}
	}
	return coordinateList
}

// abs takes the absolute value of an integer. Math.Abs is not used since that
// requires casting to a float64.
func abs(in int) int {
	if in >= 0 {
		return in
	}

	return in * -1
}

func findDistance(first *coordinate, second *coordinate) int {
	xDiff := abs(second.col - first.col)
	yDiff := abs(second.row - first.row)
	return xDiff + yDiff
}

func sumDistances(coordinateList []*coordinate) int {
	sum := 0
	for i, galaxy := range coordinateList {
		for j := i + 1; j < len(coordinateList); j++ {
			nextGalaxy := coordinateList[j]
			sum += findDistance(galaxy, nextGalaxy)
		}
	}

	return sum
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	galaxyLocations := parse(fileLines)

	fmt.Printf("The sum of all distance pairs is %d.\n", sumDistances(galaxyLocations))
}
