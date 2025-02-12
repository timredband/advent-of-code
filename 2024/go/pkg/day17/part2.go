package day17

import (
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// couldn't figure this one out. credit here https://www.youtube.com/watch?v=y-UPxMAh2N8.
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	r := regexp.MustCompile("[0-9]+")

	program := make([]int, 0)

	nums := r.FindAllString(input[3], -1)
	for i := range nums {
		num, _ := strconv.Atoi(nums[i])
		program = append(program, num)
	}

	result := find(program, 0)

	return result
}

func find(program []int, ans int) int {
	if len(program) == 0 {
		return ans
	}

	for i := range 8 {
		a := ans<<3 | i
		b := a % 8
		b = b ^ 1
		c := a / (int(math.Pow(2, float64(b))))
		b = b ^ c
		b = b ^ 4
		if b%8 == program[len(program)-1] {
			sub := find(program[:len(program)-1], a)

			if sub == 0 {
				continue
			}

			return sub
		}
	}

	return 0
}
