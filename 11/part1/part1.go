package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

type coordinate struct {
	row int
	col int
}

func parse(fileLines []string) []*coordinate {
	coordinateList := make([]*coordinate, 0)
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
			row++
		}
		row++
	}

	// Iterate through the assembled coordinates and update column values.
	for _, coord := range coordinateList {
		originalColumn := coord.col
		for _, blankColumn := range blankColumnList {
			if originalColumn >= blankColumn {
				coord.col++
			}
		}
	}
	return coordinateList
}

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
