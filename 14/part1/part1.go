package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

func calculateMatrix(matrix []string) int {
	load := 0
	for j := range matrix[0] {
		weight := len(matrix)
		for i := 0; i < len(matrix); i++ {
			char := matrix[i][j]

			if char == '#' {
				weight = len(matrix) - (i + 1)
				continue
			}

			if char == 'O' {
				load += weight
				weight--
			}
		}
	}

	return load
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	load := calculateMatrix(fileLines)

	fmt.Printf("The total load on the north support beams is %d.\n", load)
}
