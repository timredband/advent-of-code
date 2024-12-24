package day9

import (
	"os"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)[0]

	memory := make([]Memory, 0)
	freeMemory := make([]Memory, 0)

	fileId := 0

	position := 0

	for i := 0; i < len(input); i += 2 {
		fileBlock := string(input[i])
		spaceBlock := ""

		if i+1 != len(input) {
			spaceBlock = string(input[i+1])
		}

		memoryLength, _ := strconv.Atoi(fileBlock)
		freeMemoryLength, _ := strconv.Atoi(spaceBlock)

		memory = append(memory, Memory{Id: fileId, Length: memoryLength, Position: position})
		position += memoryLength

		if freeMemoryLength != 0 {
			freeMemory = append(freeMemory, Memory{Id: -1, Length: freeMemoryLength, Position: position})
			position += freeMemoryLength
		}

		fileId += 1
	}

	for j := len(memory) - 1; j >= 0; j -= 1 {
		for i := range freeMemory {
			if freeMemory[i].Position < memory[j].Position && freeMemory[i].Length >= memory[j].Length {
				memory[j].Position = freeMemory[i].Position

				freeMemory[i].Position += memory[j].Length
				freeMemory[i].Length -= memory[j].Length

				break
			}

		}
	}

	result := 0

	for _, m := range memory {
		for j := 0; j < m.Length; j += 1 {
			result += ((m.Position + j) * m.Id)
		}
	}

	return result
}
