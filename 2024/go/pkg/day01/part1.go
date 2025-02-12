package day01

import (
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	inputs := utils.ReadFile(file)

	left := make([]int, 0, len(inputs))
	right := make([]int, 0, len(inputs))

	for index, value := range inputs {
		r, _ := regexp.Compile("[0-9]+")
		input := r.FindAllString(value, -1)

		leftInt, _ := strconv.Atoi(input[0])
		rightInt, _ := strconv.Atoi(input[1])

		left = append(left, leftInt)
		right = append(right, leftInt)

		left[index] = leftInt
		right[index] = rightInt
	}

	sort.Ints(left)
	sort.Ints(right)

	result := 0

	for i := 0; i < len(left); i += 1 {
		diff := left[i] - right[i]

		if diff < 0 {
			diff *= -1
		}

		result += diff
	}

	return result
}
