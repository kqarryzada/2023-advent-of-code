package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
)

func isDigit(input rune) bool {
	return input >= '0' && input <= '9'
}

func isSpecialCharacter(input rune) bool {
	return input != '.' && !isDigit(input)
}

func getValueAndOverwrite(row int, column int, matrix [][]rune) int {
	retval := int(matrix[row][column]) - '0'
	matrix[row][column] = '.'
	return retval
}

// extractNumericalValue obtains a number from the string matrix. For example,
// for the following matrix:
//
// ...
// 4.*
// .12
// ..*
//
// If the function is provided with a row and column value of (2, 1), the number
// "12" will be returned. The matrix will also be modified to replace the '1'
// and '2' characters with '.' so that the value will not be doubly counted in a
// future invocation.
//
// If the function is provided with a row and column value that does not point
// to a digit (e.g., (0, 0)), the function will return 0.
func extractNumericalValue(row int, column int, matrix [][]rune) int {
	if !isDigit(matrix[row][column]) {
		return 0
	}

	retval := getValueAndOverwrite(row, column, matrix)

	// Iterate through values on the left and update the calculated number.
	base := 10
	for i := column - 1; i >= 0; i-- {
		if !isDigit(matrix[row][i]) {
			break
		}

		value := getValueAndOverwrite(row, i, matrix)
		retval += base * value

		base *= 10
	}

	lastIndex := len(matrix[0]) - 1
	for i := column + 1; i <= lastIndex; i++ {
		if !isDigit(matrix[row][i]) {
			break
		}

		value := getValueAndOverwrite(row, i, matrix)
		retval = (retval * 10) + value
	}

	return retval
}

// calculateLocalSum takes in the coordinates of a spceial character and fetches
// the sum of its neighboring numbers in the matrix, if any exist. The contents
// of the matrix will be updated so that a number cannot be included in the
// overall sum twice.
func calculateLocalSum(row int, column int, matrix [][]rune) int {
	localSum := 0
	lastColumnIndex := len(matrix[0]) - 1
	lastRowIndex := len(matrix) - 1

	// In most cases, the local sum is calculated by starting at
	// [row - 1, column -1] in the matrix and checking for digits from left to
	// right, with nine checks in total. The values of i and j are bounded so
	// that they never exceed the boundaries of the matrix.
	startingRow := max(row-1, 0)
	endingRow := min(row+1, lastRowIndex)
	startingColumn := max(column-1, 0)
	endingColumn := min(column+1, lastColumnIndex)
	for i := startingRow; i <= endingRow; i++ {
		for j := startingColumn; j <= endingColumn; j++ {
			localSum += extractNumericalValue(i, j, matrix)
		}
	}

	return localSum
}

func processRow(rowNumber int, matrix [][]rune) int {
	rowSum := 0
	for i, char := range matrix[rowNumber] {
		if isSpecialCharacter(char) {
			rowSum += calculateLocalSum(rowNumber, i, matrix)
		}
	}

	return rowSum
}

func processMatrix(matrix [][]rune) int {
	sum := 0
	for row := 0; row < len(matrix); row++ {
		sum += processRow(row, matrix)
	}

	return sum
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")
	matrix := make([][]rune, 0)
	for _, line := range fileLines {
		charArray := []rune(line)
		matrix = append(matrix, charArray)
	}

	sum := processMatrix(matrix)
	fmt.Printf("The sum of all the part numbers is %d.\n", sum)
}
