package day03

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
		r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
		input := r.FindAllString(line, -1)
		r, _ = regexp.Compile("[0-9]+")

		for i := 0; i < len(input); i += 1 {
			nums := r.FindAllString(input[i], -1)
			first, _ := strconv.Atoi(nums[0])
			second, _ := strconv.Atoi(nums[1])
			result += (first * second)
		}

	}

	return result
}
