package day04

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)

	check := func(i int, j int) int {
		if i-1 < 0 {
			return 0
		}

		if j-1 < 0 {
			return 0
		}

		if i+1 > len(inputs)-1 {
			return 0
		}

		if j+1 > len(inputs[i])-1 {
			return 0
		}

		diagonalLeft := (string(inputs[i-1][j-1]) == "M" && string(inputs[i+1][j+1]) == "S") || (string(inputs[i-1][j-1]) == "S" && string(inputs[i+1][j+1]) == "M")
		diagonalRight := (string(inputs[i+1][j-1]) == "M" && string(inputs[i-1][j+1]) == "S") || (string(inputs[i+1][j-1]) == "S" && string(inputs[i-1][j+1]) == "M")

		if diagonalLeft && diagonalRight {
			return 1
		}

		return 0
	}

	result := 0

	for i := 0; i < len(inputs); i += 1 {
		for j := 0; j < len(inputs[i]); j += 1 {
			if string(inputs[i][j]) == "A" {
				result += check(i, j)
			}
		}
	}

	return result
}
