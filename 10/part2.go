package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"slices"
)

type node struct {
	nodeType metalPieceType
	distance int

	// Coordinate in the grid.
	row int
	col int
}

// This represents the graph form of the input file as a matrix of node objects.
var graph [][](*node)
var NUM_ROWS int
var NUM_COLUMNS int

// metalPieceType is the type of a node in the grid of metal pieces.
type metalPieceType int

const (
	period metalPieceType = iota
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

// fetchNorthNeighbor safely fetches the north neighbor, or returns nil if the
// neighbor is unconnected or out-of-bounds.
func fetchNorthNeighbor(inputNode *node) *node {
	row := inputNode.row - 1
	col := inputNode.col
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalPieceType{
		period,
		dash,
		elbowF,
		elbow7,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	northNode := graph[row][col]
	invalidNeighbors := []metalPieceType{
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

// fetchSouthNeighbor safely fetches the south neighbor, or returns nil if the
// neighbor is unconnected or out-of-bounds.
func fetchSouthNeighbor(inputNode *node) *node {
	row := inputNode.row + 1
	col := inputNode.col
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalPieceType{
		period,
		dash,
		elbowL,
		elbowJ,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	southNode := graph[row][col]
	invalidNeighbors := []metalPieceType{
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

// fetchWestNeighbor safely fetches the west neighbor, or returns nil if the
// neighbor is unconnected or out-of-bounds.
func fetchWestNeighbor(inputNode *node) *node {
	row := inputNode.row
	col := inputNode.col - 1
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalPieceType{
		period,
		pipe,
		elbowL,
		elbowF,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	westNode := graph[row][col]
	invalidNeighbors := []metalPieceType{
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

// fetchEastNeighbor safely fetches the east neighbor, or returns nil if the
// neighbor is unconnected or out-of-bounds.
func fetchEastNeighbor(inputNode *node) *node {
	row := inputNode.row
	col := inputNode.col + 1
	if isOutOfBounds(row, col) {
		return nil
	}

	invalidTypes := []metalPieceType{
		period,
		pipe,
		elbowJ,
		elbow7,
	}
	if slices.Contains(invalidTypes, inputNode.nodeType) {
		return nil
	}

	eastNode := graph[row][col]
	invalidNeighbors := []metalPieceType{
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

// appendToQueueWithDistance updates the distance value of a node and places it
// in the provided queue. This function will safely handle nil or
// previously-visited input nodes by returning the provided queue.
func appendToQueueWithDistance(queue []*node, inputNode *node, distance int) []*node {
	if inputNode == nil {
		return queue
	}

	if inputNode.distance != 0 || inputNode.nodeType == starting {
		return queue
	}

	inputNode.distance = distance
	return append(queue, inputNode)
}

// This function finds the loop contained within the graph. For each node in the
// loop, this function computes its distance from the starting point. The
// largest value seen during this calculation is returned.
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
			queue = appendToQueueWithDistance(queue, neighbor, distance)
		}
	}

	// The final iteration in the loop will increment the maxDistance, but that
	// node will have no unvisited neighbors.
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

	var nodeType metalPieceType
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

// getNodeType fetches the 'nodeType' field of a node in the graph. This
// function computes the value of the "S" character in the grid, so
// 'starting' will never be returned. This allows callers to avoid the ambiguity
// of the "S" character by instead using the true value of "S".
func getNodeType(inputNode *node) metalPieceType {
	if inputNode.nodeType != starting {
		return inputNode.nodeType
	}

	// Infer the underlying value of the starting node from the value of the
	// neighboring nodes.
	north := fetchNorthNeighbor(inputNode)
	south := fetchSouthNeighbor(inputNode)
	west := fetchWestNeighbor(inputNode)
	east := fetchEastNeighbor(inputNode)

	if north != nil && south != nil {
		return pipe
	}
	if north != nil && west != nil {
		return elbowJ
	}
	if north != nil && east != nil {
		return elbowL
	}
	if south != nil && west != nil {
		return elbow7
	}
	if south != nil && east != nil {
		return elbowF
	}
	if west != nil && east != nil {
		return dash
	}

	panic("The starting character did not join two fields, which is not expected from the input file.")
}

// isLoopMember returns true if the provided input node is part of the closed
// loop from the input file.
func isLoopMember(inputNode *node) bool {
	return inputNode.nodeType == starting || inputNode.distance > 0
}

func findEnclosedValueCount() int {
	enclosedCount := 0

	for i := range graph {
		enclosed := false

		for j := 0; j < len(graph[0]); j++ {
			currentNode := graph[i][j]

			if !isLoopMember(currentNode) {
				if enclosed {
					enclosedCount++
				}
				continue
			}

			nodeType := getNodeType(currentNode)
			if nodeType == pipe {
				enclosed = !enclosed
				continue
			}

			// Values in the loop that reach here will have an elbow nodeType.
			// Iterate until we skip over all subsequent dash characters in this
			// row.
			prevNodeType := nodeType
			for j++; j < len(graph[0]); j++ {
				if getNodeType(graph[i][j]) != dash {
					break
				}
			}

			nodeType = getNodeType(graph[i][j])
			if prevNodeType == elbowL && nodeType == elbowJ ||
				prevNodeType == elbowF && nodeType == elbow7 {

				// We've iterated through a horizontal bottom/top piece, such as
				// the bottom line of this example:
				// |    |
				// L----J
				//
				// Therefore, the 'enclosed' status has not changed.
				continue
			}

			enclosed = !enclosed
		}
	}

	return enclosedCount
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	initializeGraph(&fileLines)
	computeGraph(&fileLines)
	fmt.Printf("The number of enclosed tiles is %d.\n", findEnclosedValueCount())
}
