package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"log"
	"strconv"
	"strings"
)

type gameSet struct {
	red   int
	green int
	blue  int
}

// parseGameData parses the string representation of a game history record into
// a list of gameSet objects. An example history record can take the form of:
//
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGameData(gameData string) []gameSet {
	gameList := make([]gameSet, 0)

	// Obtain the string data after "Game n: ".
	dataString := strings.Split(gameData, ": ")[1]

	subsets := strings.Split(dataString, "; ")
	for _, set := range subsets {
		record := new(gameSet)

		colorCounts := strings.Split(set, ", ")
		// colorCount is an individual record that takes the form of "3 blue".
		for _, colorCount := range colorCounts {
			countString := strings.Split(colorCount, " ")[0]
			count, err := strconv.Atoi(countString)
			if err != nil {
				log.Fatal("An unexpected error occurred during parsing.", err)
			}

			if strings.Contains(colorCount, "red") {
				record.red = count
			} else if strings.Contains(colorCount, "green") {
				record.green = count
			} else if strings.Contains(colorCount, "blue") {
				record.blue = count
			} else {
				log.Fatal("The game record did not contain an expected color.", gameData)
			}
		}

		gameList = append(gameList, *record)
	}

	return gameList
}

func main() {
	sum := 0

	// This object represents the maximum constraints for the game.
	maximum := new(gameSet)
	maximum.red = 12
	maximum.green = 13
	maximum.blue = 14

	fileLines := fileutils.LoadFile("input.txt")
	for lineNumber, line := range fileLines {
		gameNumber := lineNumber + 1
		gameDataList := parseGameData(line)

		possible := true
		for _, record := range gameDataList {
			if record.red > maximum.red ||
				record.green > maximum.green ||
				record.blue > maximum.blue {

				possible = false
				break
			}
		}

		if possible {
			sum += gameNumber
		}
	}

	fmt.Printf("The total sum of the possible game numbers is %d.\n", sum)
}
