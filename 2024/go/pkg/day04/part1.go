package day04

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	inputs := utils.ReadFile(file)

	checkDiagonalUpLeft := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i-3 < 0 {
			return false
		}

		if j-3 < 0 {
			return false
		}

		if string(inputs[i-1][j-1]) == "M" && string(inputs[i-2][j-2]) == "A" && string(inputs[i-3][j-3]) == "S" {
			return true
		}

		return false
	}

	checkUp := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i-3 < 0 {
			return false
		}

		if string(inputs[i-1][j]) == "M" && string(inputs[i-2][j]) == "A" && string(inputs[i-3][j]) == "S" {
			return true
		}

		return false
	}

	checkDiagonalUpRight := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i-3 < 0 {
			return false
		}

		if j+3 > len(inputs[i])-1 {
			return false
		}

		if string(inputs[i-1][j+1]) == "M" && string(inputs[i-2][j+2]) == "A" && string(inputs[i-3][j+3]) == "S" {
			return true
		}

		return false
	}

	checkRight := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if j+3 > len(inputs[i])-1 {
			return false
		}

		if string(inputs[i][j+1]) == "M" && string(inputs[i][j+2]) == "A" && string(inputs[i][j+3]) == "S" {
			return true
		}

		return false
	}

	checkDiagonalDownRight := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i+3 > len(inputs)-1 {
			return false
		}

		if j+3 > len(inputs[i])-1 {
			return false
		}

		if string(inputs[i+1][j+1]) == "M" && string(inputs[i+2][j+2]) == "A" && string(inputs[i+3][j+3]) == "S" {
			return true
		}

		return false
	}

	checkDown := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i+3 > len(inputs)-1 {
			return false
		}

		if string(inputs[i+1][j]) == "M" && string(inputs[i+2][j]) == "A" && string(inputs[i+3][j]) == "S" {
			return true
		}

		return false
	}

	checkDiagonalDownLeft := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if i+3 > len(inputs)-1 {
			return false
		}

		if j-3 < 0 {
			return false
		}

		if string(inputs[i+1][j-1]) == "M" && string(inputs[i+2][j-2]) == "A" && string(inputs[i+3][j-3]) == "S" {
			return true
		}

		return false
	}

	checkLeft := func(i int, j int) bool {
		if string(inputs[i][j]) != "X" {
			return false
		}

		if j-3 < 0 {
			return false
		}

		if string(inputs[i][j-1]) == "M" && string(inputs[i][j-2]) == "A" && string(inputs[i][j-3]) == "S" {
			return true
		}

		return false
	}

	find := func(i int, j int) int {
		count := 0

		if checkDiagonalUpLeft(i, j) {
			count += 1
		}

		if checkUp(i, j) {
			count += 1
		}

		if checkDiagonalUpRight(i, j) {
			count += 1
		}

		if checkRight(i, j) {
			count += 1
		}

		if checkDiagonalDownRight(i, j) {
			count += 1
		}

		if checkDown(i, j) {
			count += 1
		}

		if checkDiagonalDownLeft(i, j) {
			count += 1
		}

		if checkLeft(i, j) {
			count += 1
		}

		return count
	}

	result := 0

	for i := 0; i < len(inputs); i += 1 {
		for j := 0; j < len(inputs[i]); j += 1 {
			if string(inputs[i][j]) == "X" {
				result += find(i, j)
			}
		}
	}

	return result
}
