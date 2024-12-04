package day2

import (
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
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

		if integers[0] == integers[1] {
			continue
		}

		increasing := integers[0] < integers[1]
		isSafe := true

		for i := range integers {
			if i == 0 {
				continue
			}

			if integers[i] == integers[i-1] {
				isSafe = false
				break
			}

			if (increasing && integers[i-1] > integers[i]) || (!increasing && integers[i-1] < integers[i]) {
				isSafe = false
				break
			}

			diff := integers[i] - integers[i-1]

			if diff < 0 {
				diff *= -1
			}

			if diff > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			result += 1
		}
	}

	return result
}
