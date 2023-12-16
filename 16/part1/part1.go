package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

var NUM_ROWS int
var NUM_COLS int

type tile struct {
	Type        tileType
	isEnergized bool

	// Tracks whether the tile has had light travel through it in a particular
	// direction. This provides a quick escape if light has already travelled here
	// before, and prevents getting stuck if the light travels in a loop.
	hasUpTravelled    bool
	hasDownTravelled  bool
	hasLeftTravelled  bool
	hasRightTravelled bool
}

type tileType int

const (
	EMPTY_SPACE tileType = iota
	MIRROR_FORWARD_SLASH
	MIRROR_BACKSLASH
	SPLIT_HORIZTONAL
	SPLIT_VERTICAL
)

type direction int

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

func parse(fileLines []string) [][]*tile {
	grid := make([][]*tile, 0)

	for _, line := range fileLines {
		gridLine := make([]*tile, 0)
		for _, char := range line {
			newTile := new(tile)
			var newType tileType

			switch char {
			case '.':
				newType = EMPTY_SPACE
			case '/':
				newType = MIRROR_FORWARD_SLASH
			case '\\':
				newType = MIRROR_BACKSLASH
			case '-':
				newType = SPLIT_HORIZTONAL
			case '|':
				newType = SPLIT_VERTICAL
			default:
				panic("Invalid tile type found: " + string(char))
			}

			newTile.Type = newType
			gridLine = append(gridLine, newTile)
		}

		grid = append(grid, gridLine)
	}

	NUM_ROWS = len(grid)
	NUM_COLS = len(grid[0])

	return grid
}

func isOutOfBounds(row int, col int) bool {
	if row < 0 || row >= NUM_ROWS {
		return true
	}

	return col < 0 || col >= NUM_COLS
}

func nextTile(row int, col int, dir direction) (int, int) {
	newRow := row
	newCol := col

	if dir == LEFT {
		newCol--
	} else if dir == RIGHT {
		newCol++
	} else if dir == UP {
		newRow--
	} else if dir == DOWN {
		newRow++
	}

	return newRow, newCol
}

func nextDirectionForMirror(originalDirection direction, tile tileType) direction {
	if tile != MIRROR_BACKSLASH && tile != MIRROR_FORWARD_SLASH {
		panic("Improper mirror type entered.")
	}

	if tile == MIRROR_FORWARD_SLASH {
		switch originalDirection {
		case UP:
			return RIGHT
		case DOWN:
			return LEFT
		case RIGHT:
			return UP
		case LEFT:
			return DOWN
		default:
			panic("Invalid direction entered.")
		}
	}

	switch originalDirection {
	case UP:
		return LEFT
	case DOWN:
		return RIGHT
	case RIGHT:
		return DOWN
	case LEFT:
		return UP
	default:
		panic("Invalid direction entered.")
	}
}

// followPath recursively traverses the provided grid given an initial location
// and direction.
func followPath(row int, col int, grid [][]*tile, dir direction) {
	if isOutOfBounds(row, col) {
		return
	}

	originalDirection := dir
	myTile := grid[row][col]
	myTile.isEnergized = true

	// Check if we have previously passed through this tile in this direction
	// already.
	switch dir {
	case UP:
		if myTile.hasUpTravelled {
			return
		}
		myTile.hasUpTravelled = true
	case DOWN:
		if myTile.hasDownTravelled {
			return
		}
		myTile.hasDownTravelled = true
	case LEFT:
		if myTile.hasLeftTravelled {
			return
		}
		myTile.hasLeftTravelled = true
	case RIGHT:
		if myTile.hasRightTravelled {
			return
		}
		myTile.hasRightTravelled = true
	}

	switch myTile.Type {
	case EMPTY_SPACE:
		nextRow, nextCol := nextTile(row, col, originalDirection)
		followPath(nextRow, nextCol, grid, originalDirection)

	case MIRROR_FORWARD_SLASH:
		newDirection := nextDirectionForMirror(originalDirection, MIRROR_FORWARD_SLASH)
		newRow, newCol := nextTile(row, col, newDirection)
		followPath(newRow, newCol, grid, newDirection)

	case MIRROR_BACKSLASH:
		newDirection := nextDirectionForMirror(originalDirection, MIRROR_BACKSLASH)
		newRow, newCol := nextTile(row, col, newDirection)
		followPath(newRow, newCol, grid, newDirection)

	case SPLIT_HORIZTONAL:
		if originalDirection == LEFT {
			followPath(row, col-1, grid, LEFT)
		} else if originalDirection == RIGHT {
			followPath(row, col+1, grid, RIGHT)
		} else {
			followPath(row, col-1, grid, LEFT)
			followPath(row, col+1, grid, RIGHT)
		}

	case SPLIT_VERTICAL:
		if originalDirection == UP {
			followPath(row-1, col, grid, UP)
		} else if originalDirection == DOWN {
			followPath(row+1, col, grid, DOWN)
		} else {
			followPath(row-1, col, grid, UP)
			followPath(row+1, col, grid, DOWN)
		}

	default:
		panic("Invalid tile type found.")
	}
}

func main() {
	fileLines := utils.LoadFile("input.txt")
	grid := parse(fileLines)
	followPath(0, 0, grid, RIGHT)

	sum := 0
	for _, gridLine := range grid {
		for _, tile := range gridLine {
			if tile.isEnergized {
				sum++
			}
		}
	}

	fmt.Printf("The total number of energized tiles is %d.\n", sum)
}
