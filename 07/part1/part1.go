package main

import (
	"fmt"
	fileutils "kqarryzada/advent-of-code-2023/utils"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	raw      string
	typeName handType
	bid      int
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func containsValue(m map[byte]int, v int) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}

	return false
}

func compute(rawHand string) hand {
	if len(rawHand) != 5 {
		panic("Improper input: " + rawHand)
	}

	retval := new(hand)
	retval.raw = rawHand

	mapVals := make(map[byte]int)
	mapVals[rawHand[0]] = 1
	for i := 1; i < len(rawHand); i++ {
		char := rawHand[i]
		mapVals[char]++
	}

	if len(mapVals) == 1 {
		retval.typeName = fiveOfAKind
		return *retval
	}
	if len(mapVals) == 5 {
		retval.typeName = highCard
		return *retval
	}
	if containsValue(mapVals, 4) {
		retval.typeName = fourOfAKind
		return *retval
	}
	if containsValue(mapVals, 3) {
		if len(mapVals) == 2 {
			retval.typeName = fullHouse
		} else {
			retval.typeName = threeOfAKind
		}

		return *retval
	}
	if len(mapVals) == 3 {
		retval.typeName = twoPair
		return *retval
	}

	retval.typeName = onePair
	return *retval
}

func convertCardToInt(char byte) int {
	var num int
	if char >= '2' && char <= '9' {
		num = int(char - '0')
	} else if char == 'T' {
		num = 10
	} else if char == 'J' {
		num = 11
	} else if char == 'Q' {
		num = 12
	} else if char == 'K' {
		num = 13
	} else if char == 'A' {
		num = 14
	} else {
		panic("Invalid character found: " + string(char))
	}

	return num
}

func compareHands(hand1 hand, hand2 hand) bool {
	if hand1.typeName != hand2.typeName {
		return hand1.typeName < hand2.typeName
	}

	for i := range hand1.raw {
		val1 := convertCardToInt(hand1.raw[i])
		val2 := convertCardToInt(hand2.raw[i])

		if val1 != val2 {
			return val1 < val2
		}
	}

	panic("Two identical hands were found.")
}

func main() {
	fileLines := fileutils.LoadFile("input.txt")

	handList := make([]hand, 0)
	for _, line := range fileLines {
		values := strings.Fields(line)
		parsedHand := compute(values[0])
		parsedHand.bid, _ = strconv.Atoi(values[1])

		handList = append(handList, parsedHand)
	}

	sort.Slice(handList, func(i, j int) bool {
		return compareHands(handList[i], handList[j])
	})

	totalWinnings := 0
	for i, handEntry := range handList {
		totalWinnings += (i + 1) * handEntry.bid
	}

	fmt.Printf("The total winnings across all the hands are %d.\n", totalWinnings)
}
