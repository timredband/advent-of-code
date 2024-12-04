package day1

import (
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)

	left := make([]int, 0, len(inputs))
	right := make(map[int]int)

	for _, value := range inputs {
		r, _ := regexp.Compile("[0-9]+")
		input := r.FindAllString(value, -1)

		leftInt, _ := strconv.Atoi(input[0])
		rightInt, _ := strconv.Atoi(input[1])

		left = append(left, leftInt)
		right[rightInt] += 1
	}

	result := 0

	for i := 0; i < len(left); i += 1 {
		similarity := left[i] * right[left[i]]
		result += similarity
	}

	return result
}
