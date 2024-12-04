package day2

import (
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)
	result := 0

	for _, line := range inputs {
		r, _ := regexp.Compile("[0-9]+")
		input := r.FindAllString(line, -1)

		integers := make([]int, 0, len(inputs))

		for _, v := range input {
			integer, _ := strconv.Atoi(v)
			integers = append(integers, integer)
		}

		isSafe := isRowSafe(integers)

		if !isSafe {
			for i := range integers {
				foo := make([]int, len(integers))
				copy(foo, integers)

				foo = append(foo[:i], foo[i+1:]...)

				isSafe = isRowSafe(foo)
				if isSafe {
					break
				}
			}

		}

		if isSafe {
			result += 1
		}
	}

	return result
}

func isRowSafe(row []int) bool {
	increasing := row[0] < row[1]
	isSafe := true

	for i := range row {
		if i == 0 {
			continue
		}

		if row[i] == row[i-1] {
			isSafe = false
			break
		}

		if (increasing && row[i-1] > row[i]) || (!increasing && row[i-1] < row[i]) {
			isSafe = false
			break
		}

		diff := row[i] - row[i-1]

		if diff < 0 {
			diff *= -1
		}

		if diff > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}
