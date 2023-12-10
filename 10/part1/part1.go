package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"slices"
)

type node struct {
	nodeType metalJoin
	distance int

	// Coordinate in the grid.
	row int
	col int
}

var graph [][](*node)
var NUM_ROWS int
var NUM_COLUMNS int

type metalJoin int

const (
	period metalJoin = iota
	pipe
	dash
	elbowL
	elbowJ
	elbowF
	elbow7
	starting
)

func isOutOfBounds(row int, col int) bool {
	if row < 0 || row > NUM_ROWS {
		return true
	}
	if col < 0 || col > NUM_COLUMNS {
		return true
	}

	return false
}

func fetchNorthNeighbor(inputNode *node) *node {
	row := inputNode.row - 1
	col := inputNode.col
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalJoin{
		period,
		dash,
		elbowF,
		elbow7,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	northNode := graph[row][col]
	invalidNeighbors := []metalJoin{
		period,
		dash,
		elbowL,
		elbowJ,
		starting,
	}
	if slices.Contains(invalidNeighbors, northNode.nodeType) {
		return nil
	}

	return northNode

}

func fetchSouthNeighbor(inputNode *node) *node {
	row := inputNode.row + 1
	col := inputNode.col
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalJoin{
		period,
		dash,
		elbowL,
		elbowJ,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	southNode := graph[row][col]
	invalidNeighbors := []metalJoin{
		period,
		dash,
		elbowF,
		elbow7,
		starting,
	}
	if slices.Contains(invalidNeighbors, southNode.nodeType) {
		return nil
	}

	return southNode
}

func fetchWestNeighbor(inputNode *node) *node {
	row := inputNode.row
	col := inputNode.col - 1
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalJoin{
		period,
		pipe,
		elbowL,
		elbowF,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	westNode := graph[row][col]
	invalidNeighbors := []metalJoin{
		period,
		pipe,
		elbowJ,
		elbow7,
		starting,
	}
	if slices.Contains(invalidNeighbors, westNode.nodeType) {
		return nil
	}

	return westNode
}

func fetchEastNeighbor(inputNode *node) *node {
	row := inputNode.row
	col := inputNode.col + 1
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalJoin{
		period,
		pipe,
		elbowJ,
		elbow7,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	eastNode := graph[row][col]
	invalidNeighbors := []metalJoin{
		period,
		pipe,
		elbowL,
		elbowF,
		starting,
	}
	if slices.Contains(invalidNeighbors, eastNode.nodeType) {
		return nil
	}

	return eastNode
}

func appendToQueue(queue []*node, inputNode *node, distance int) []*node {
	if inputNode == nil {
		return queue
	}

	// Check if the node has already been visited.
	if inputNode.distance != 0 || inputNode.nodeType == starting {
		return queue
	}

	inputNode.distance = distance
	return append(queue, inputNode)
}

func computeGraph(inputLines *[]string) int {
	lineNumber, charIndex := findStartingPoint()
	startingNode := graph[lineNumber][charIndex]
	startingNode.distance = 0

	// Initialize the queue. Use the standard append() function since
	// appendToQueue() avoids adding the starting node.
	queue := make([]*node, 0)
	queue = append(queue, startingNode)

	maxDistance := 0
	var neighbors [4](*node)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		distance := currentNode.distance + 1
		maxDistance = max(maxDistance, distance)

		neighbors[0] = fetchNorthNeighbor(currentNode)
		neighbors[1] = fetchSouthNeighbor(currentNode)
		neighbors[2] = fetchWestNeighbor(currentNode)
		neighbors[3] = fetchEastNeighbor(currentNode)
		for _, neighbor := range neighbors {
			queue = appendToQueue(queue, neighbor, distance)
		}
	}

	return maxDistance - 1
}

func findStartingPoint() (x int, y int) {
	for i := range graph {
		for j := range graph[0] {
			if graph[i][j].nodeType == starting {
				return i, j
			}
		}
	}

	panic("Starting character not found.")
}

func createNode(char rune, row int, col int) *node {
	returnNode := &node{
		row: row,
		col: col,
	}

	var nodeType metalJoin
	switch char {
	case '.':
		nodeType = period
	case '|':
		nodeType = pipe
	case '-':
		nodeType = dash
	case 'L':
		nodeType = elbowL
	case 'J':
		nodeType = elbowJ
	case 'F':
		nodeType = elbowF
	case '7':
		nodeType = elbow7
	case 'S':
		nodeType = starting
	default:
		panic(fmt.Sprintf("Invalid letter found: %c", char))
	}

	returnNode.nodeType = nodeType
	return returnNode
}

func initializeGraph(inputFile *[]string) {
	graph = make([][]*node, 0)

	fileLines := *inputFile
	for i, line := range fileLines {
		gridLine := make([]*node, len(line))

		for j, char := range line {
			newNode := createNode(char, i, j)
			gridLine[j] = newNode
		}

		graph = append(graph, gridLine)
	}

	NUM_ROWS = len(graph) - 1
	NUM_COLUMNS = len(graph[0]) - 1
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	initializeGraph(&fileLines)
	maxDistance := computeGraph(&fileLines)
	fmt.Printf("The largest distance found was %d.\n", maxDistance)
}
