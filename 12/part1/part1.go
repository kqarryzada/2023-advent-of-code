package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
)

func main() {
	fileLines := utils.LoadFile("example.txt")
	println(fileLines[0][0])

	fmt.Printf("Answer: %d.\n", -1)
}
