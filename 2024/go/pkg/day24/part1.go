package day24

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type instruction struct {
	leftInputWire   string
	rightInputWire  string
	gateType        string
	destinationWire string
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	wiresByName := make(map[string]int)
	instructions := make(map[string]instruction)

	for i := range input {
		if strings.Contains(input[i], ":") {
			name := input[i][0:3]
			wiresByName[name], _ = strconv.Atoi(string(input[i][len(input[i])-1]))
			continue
		}

		r := regexp.MustCompile("[A-Za-z0-9]+")
		instructionInput := r.FindAllString(input[i], -1)

		if _, ok := wiresByName[instructionInput[0]]; !ok {
			wiresByName[instructionInput[0]] = -1
		}

		if _, ok := wiresByName[instructionInput[2]]; !ok {
			wiresByName[instructionInput[2]] = -1
		}

		if _, ok := wiresByName[instructionInput[3]]; !ok {
			wiresByName[instructionInput[3]] = -1
		}

		instruction := instruction{
			leftInputWire:   instructionInput[0],
			rightInputWire:  instructionInput[2],
			gateType:        instructionInput[1],
			destinationWire: instructionInput[3],
		}

		instructions[instructionInput[3]] = instruction
	}

	for len(instructions) != 0 {
		for destinationWire, instruction := range instructions {
			if wiresByName[instruction.leftInputWire] != -1 && wiresByName[instruction.rightInputWire] != -1 {
				leftValue := wiresByName[instruction.leftInputWire]
				rightValue := wiresByName[instruction.rightInputWire]

				switch instruction.gateType {
				case "AND":
					wiresByName[destinationWire] = leftValue & rightValue
				case "OR":
					wiresByName[destinationWire] = leftValue | rightValue
				case "XOR":
					wiresByName[destinationWire] = leftValue ^ rightValue
				}

				delete(instructions, destinationWire)
			}
		}
	}

	i := 0
	result := 0

	for {
		key := fmt.Sprintf("z%02d", i)

		if value, ok := wiresByName[key]; ok {
			result += (value * int(math.Pow(2, float64(i))))
			i += 1
		} else {
			break
		}
	}

	return result
}
