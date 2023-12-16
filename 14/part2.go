package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

// The number of cycles that were requested.
var NUM_CYCLES int = 1_000_000_000

// The number of cycles that should be performed before checking for a repeating
// sequence of matrix cycles.
var WINDUP_CYCLE_COUNT = 100

// The bounds of the matrix.
var NUM_ROWS int
var NUM_COLS int

func slideNorth(inputMatrix *[][]rune) {
	matrix := *inputMatrix

	for j := range matrix[0] {
		nextSlot := 0
		for i := 0; i < NUM_ROWS; i++ {
			char := matrix[i][j]

			if char == '#' {
				nextSlot = i + 1
				continue
			}

			if char == 'O' {
				// Swap the two locations if the 'O' character is not already
				// in the correct position.
				if nextSlot != i {
					matrix[nextSlot][j] = 'O'
					matrix[i][j] = '.'
				}
				nextSlot = min(nextSlot+1, NUM_ROWS-1)
			}
		}
	}
}

func slideSouth(inputMatrix *[][]rune) {
	matrix := *inputMatrix

	for j := range matrix[0] {
		nextSlot := NUM_ROWS - 1
		for i := (NUM_ROWS - 1); i >= 0; i-- {
			char := matrix[i][j]

			if char == '#' {
				nextSlot = i - 1
				continue
			}

			if char == 'O' {
				if nextSlot != i {
					matrix[nextSlot][j] = 'O'
					matrix[i][j] = '.'
				}
				nextSlot = max(nextSlot-1, 0)
			}
		}
	}
}

func slideWest(inputMatrix *[][]rune) {
	matrix := *inputMatrix

	for i := range matrix {
		nextSlot := 0
		for j := 0; j < NUM_COLS; j++ {
			char := matrix[i][j]

			if char == '#' {
				nextSlot = j + 1
				continue
			}

			if char == 'O' {
				if nextSlot != j {
					matrix[i][nextSlot] = 'O'
					matrix[i][j] = '.'
				}
				nextSlot = min(nextSlot+1, NUM_COLS-1)
			}
		}
	}
}

func slideEast(inputMatrix *[][]rune) {
	matrix := *inputMatrix

	for i := range matrix {
		nextSlot := NUM_COLS - 1
		for j := (NUM_COLS - 1); j >= 0; j-- {
			char := matrix[i][j]

			if char == '#' {
				nextSlot = j - 1
				continue
			}

			if char == 'O' {
				if nextSlot != j {
					matrix[i][nextSlot] = 'O'
					matrix[i][j] = '.'
				}
				nextSlot = max(nextSlot-1, 0)
			}
		}
	}
}

// cycle performs a full iteration of sliding the 'O' values around the matrix.
func cycle(matrix *[][]rune) {
	slideNorth(matrix)
	slideWest(matrix)
	slideSouth(matrix)
	slideEast(matrix)
}

func convertToMatrix(stringLines []string) *[][]rune {
	matrix := make([][]rune, 0)
	for _, line := range stringLines {
		array := []rune(line)
		matrix = append(matrix, array)
	}

	NUM_ROWS = len(matrix)
	NUM_COLS = len(matrix[0])

	return &matrix
}

// For a given state in the matrix, calculateMatrixLoad computes the numerical
// "load" of the matrix, which is dependent on the location of 'O' characters.
func calculateMatrixLoad(inputMatrix *[][]rune) int {
	matrix := *inputMatrix
	matrixLength := len(matrix)
	load := 0
	for j := range matrix[0] {
		for i := 0; i < matrixLength; i++ {
			if matrix[i][j] == 'O' {
				load += matrixLength - i
			}
		}
	}

	return load
}

func determineCycleCount(inputMatrix *[][]rune) int {
	// Copy the input matrix to avoid modifications to the input matrix.
	origMatrix := *inputMatrix
	matrix := make([][]rune, 0)
	for i := range origMatrix {
		row := make([]rune, 0)
		for j := 0; j < len(origMatrix[0]); j++ {
			row = append(row, origMatrix[i][j])
		}
		matrix = append(matrix, row)
	}

	// Cycle the matrix for several iterations so that the operations reach a
	// repeating sequence.
	for i := 0; i < WINDUP_CYCLE_COUNT; i++ {
		cycle(&matrix)
	}

	matrixCopy := make([]string, len(matrix))
	for i, row := range matrix {
		matrixCopy[i] = string(row)
	}

	cycleCount := 0
	for {
		cycle(&matrix)
		cycleCount++

		isEqual := true
		for i, row := range matrix {
			if string(row) != matrixCopy[i] {
				isEqual = false
				break
			}
		}

		if isEqual {
			break
		} else if cycleCount >= 100_000 {
			panic("A repeating sequence was not found after 100,000 iterations. Consider increasing WINDUP_CYCLE_COUNT.")
		}
	}

	return cycleCount
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	matrix := convertToMatrix(fileLines)

	// Due to the nature of repeatedly sliding elements in 4 directions, there
	// comes a point where the sequence will repeat since there are only so many
	// possible states. Once the cycling begins to loop, we should determine the
	// number of steps required to reach the same state again.
	//
	// The purpose behind this optimization is to cut down the number of
	// iterations that this program is required to perform when the value of
	// NUM_CYCLES is very large.
	cycleCount := determineCycleCount(matrix)

	// Determining the cycle count requires some "wind up" time to allow the
	// sequence to reach a repeatable point. Wind up the matrix with the initial
	// cycle count so that future calls to cycle() are guaranteed to repeat.
	for i := 0; i < WINDUP_CYCLE_COUNT; i++ {
		cycle(matrix)
	}

	// We have already performed WINDUP_CYCLE_COUNT iterations, so this can be
	// deducted from the overall NUM_CYCLES count. The cycling should now be a
	// repeatable process.
	numIterations := (NUM_CYCLES - WINDUP_CYCLE_COUNT) % cycleCount
	for i := 0; i < numIterations; i++ {
		cycle(matrix)
	}

	fmt.Printf("The total load on the support beams is %d.\n", calculateMatrixLoad(matrix))
}
