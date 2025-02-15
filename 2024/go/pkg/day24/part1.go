package day24

import (
	"fmt"
	"github.com/timredband/advent-of-code/pkg/utils"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type instruction struct {
	leftInputWire   string
	rightInputWire  string
	gateType        string
	destinationWire string
}

func initialize(input []string) (map[string]int, map[string]instruction) {
	wiresByName := make(map[string]int)
	instructions := make(map[string]instruction)

	for i := range input {
		randomNumber := rand.Intn(2)

		if strings.Contains(input[i], ":") {
			name := input[i][0:3]
			wiresByName[name] = randomNumber
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

		if instructionInput[0] > instructionInput[2] {
			instructionInput[0], instructionInput[2] = instructionInput[2], instructionInput[0]
		}

		instruction := instruction{
			leftInputWire:   instructionInput[0],
			rightInputWire:  instructionInput[2],
			gateType:        instructionInput[1],
			destinationWire: instructionInput[3],
		}

		instructions[instructionInput[3]] = instruction
	}

	return wiresByName, instructions
}

func processInstructions(wiresByName map[string]int, instructions map[string]instruction) {
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
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	wiresByName, instructions := initialize(input)

	processInstructions(wiresByName, instructions)

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
