package main

import (
	"fmt"
	utils "kqarryzada/advent-of-code-2023/utils"
	"strings"
)

type comparator int

const (
	LESS_THAN comparator = iota
	GREATER_THAN
)

// workflowStep describes an individual step in a workflow, e.g.,
// "a<2006:qkq".
type workflowStep struct {
	paramName string
	comp      comparator
	value     int

	// i.e., "qkq"
	gotoParam string

	isGoTo bool
}

// A 'part' contains the same core 'paramName', 'comp', and 'value' fields
// of a workflowStep. The other parts are not used.
type part struct {
	x int
	m int
	a int
	s int
}

type workflow struct {
	name  string
	steps []workflowStep
}

func parseWorkflowStep(step string) workflowStep {
	wfStep := new(workflowStep)
	i := 0
	for ; i < len(step); i++ {
		char := step[i]
		if char == '<' {
			wfStep.comp = LESS_THAN
			break
		} else if char == '>' {
			wfStep.comp = GREATER_THAN
			break
		}
	}
	wfStep.paramName = step[0:i]

	remain := step[i+1:]
	remainArray := strings.Split(remain, ":")

	if len(remainArray) != 2 {
		panic("Unexpected format: " + remain)
	}
	wfStep.value = utils.AsInt(remainArray[0])
	wfStep.gotoParam = remainArray[1]

	return *wfStep
}

func parseWorkflow(rawWorkflow string) workflow {
	flow := new(workflow)
	splitString := strings.Split(rawWorkflow, "{")
	flow.name = splitString[0]

	splitString = strings.Split(splitString[1], ",")

	stepList := make([]workflowStep, 0)
	for i := 0; i < len(splitString)-1; i++ {
		step := splitString[i]
		wfStep := parseWorkflowStep(step)
		stepList = append(stepList, wfStep)
	}

	lastStep := splitString[len(splitString)-1]
	lastStep = lastStep[:len(lastStep)-1]

	lastWfStep := &workflowStep{
		isGoTo:    true,
		gotoParam: lastStep,
	}
	stepList = append(stepList, *lastWfStep)

	flow.steps = stepList
	return *flow
}

func compare(comp comparator, threshold int, value int) bool {
	switch comp {
	case LESS_THAN:
		return value < threshold
	case GREATER_THAN:
		return value > threshold
	}

	panic("Unexpected comparator found.")
}

func runWorkflow(prt part, flowMap map[string]workflow, mapEntry string) bool {
	flow := flowMap[mapEntry]
	for _, step := range flow.steps {
		if step.isGoTo {
			if step.gotoParam == "A" {
				return true
			} else if step.gotoParam == "R" {
				return false
			}

			return runWorkflow(prt, flowMap, step.gotoParam)
		}

		partValue := 0
		switch step.paramName {
		case "x":
			partValue = prt.x
		case "m":
			partValue = prt.m
		case "a":
			partValue = prt.a
		case "s":
			partValue = prt.s
		default:
			panic("Unexpected value found: " + step.paramName)
		}

		shouldGo := compare(step.comp, step.value, partValue)
		if shouldGo {
			if step.gotoParam == "A" {
				return true
			} else if step.gotoParam == "R" {
				return false
			}

			return runWorkflow(prt, flowMap, step.gotoParam)
		}
	}

	panic("The end was unexpectedly reached.")
}

func getRating(line string, flowMap map[string]workflow) int {
	// Remove braces.
	line = line[1 : len(line)-1]

	part := asPart(line)
	isAccepted := runWorkflow(part, flowMap, "in")
	if !isAccepted {
		return 0
	}

	return part.x + part.m + part.a + part.s
}

func asPart(line string) part {
	newPart := new(part)

	components := strings.Split(line, ",")
	for i, comp := range components {
		strVal := strings.Split(comp, "=")[1]

		if i == 0 {
			newPart.x = utils.AsInt(strVal)
		} else if i == 1 {
			newPart.m = utils.AsInt(strVal)
		} else if i == 2 {
			newPart.a = utils.AsInt(strVal)
		} else if i == 3 {
			newPart.s = utils.AsInt(strVal)
		}
	}

	return *newPart
}

func main() {
	fileLines := utils.LoadFile("input.txt")

	workflows := make(map[string]workflow, 0)
	i := 0
	for j, line := range fileLines {
		if len(line) == 0 {
			i = j + 1
			break
		}

		flow := parseWorkflow(line)
		workflows[flow.name] = flow
	}

	sum := 0
	for ; i < len(fileLines); i++ {
		sum += getRating(fileLines[i], workflows)
	}

	fmt.Printf("The sum of the ratings for the accepted parts is %d.\n", sum)
}
