package day9

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Memory struct {
	Id       int
	Length   int
	Position int
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)[0]

	freeMemory := Memory{Id: -1}
	memory := make([]Memory, 0)

	fileId := 0

	for i := 0; i < len(input); i += 2 {
		fileBlock := string(input[i])
		spaceBlock := ""

		if i+1 != len(input) {
			spaceBlock = string(input[i+1])
		}

		fileBlockInt, _ := strconv.Atoi(fileBlock)
		spaceBlockInt, _ := strconv.Atoi(spaceBlock)

		block := Memory{Id: fileId}

		for j := 0; j < fileBlockInt; j += 1 {
			memory = append(memory, block)
		}

		for j := 0; j < spaceBlockInt; j += 1 {
			memory = append(memory, freeMemory)
		}

		fileId += 1
	}

	i, j := 0, len(memory)-1

	for i < j {
		for i < j && memory[j] == freeMemory {
			j -= 1
		}

		for i < j && memory[i] != freeMemory {
			i += 1
		}

		memory[i], memory[j] = memory[j], memory[i]

		i += 1
		j -= 1
	}

	result := 0

	for i, v := range memory {
		if v == freeMemory {
			break
		}

		result += (i * v.Id)
	}

	return result
}
