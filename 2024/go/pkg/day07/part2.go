package day07

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)

	equations := make([]Equation, 0, len(inputs))

	for _, v := range inputs {
		line := strings.Split(v, ": ")

		target, _ := strconv.Atoi(line[0])
		inputStrings := strings.Split(line[1], " ")
		inputInts := make([]int, 0, len(inputStrings))

		for _, v := range inputStrings {
			num, _ := strconv.Atoi(v)
			inputInts = append(inputInts, num)
		}

		equation := Equation{target: target, nums: inputInts}
		equations = append(equations, equation)
	}

	var backtrack func(index int, value int, target int, nums []int) bool
	backtrack = func(index int, value int, target int, nums []int) bool {
		if index == len(nums)-1 && value == target {
			return true
		}

		if index == len(nums)-1 {
			return false
		}

		plusResult := backtrack(index+1, value+nums[index+1], target, nums)
		if plusResult {
			return true
		}

		multiResult := backtrack(index+1, value*nums[index+1], target, nums)
		if multiResult {
			return true
		}

		concatted, _ := strconv.Atoi(strconv.Itoa(value) + strconv.Itoa(nums[index+1]))
		concatResult := backtrack(index+1, concatted, target, nums)
		if concatResult {
			return true
		}

		return false
	}

	result := 0

	for _, equation := range equations {
		if len(equation.nums) == 0 {
			continue
		}

		if len(equation.nums) == 1 && equation.nums[0] == equation.target {
			result += equation.target
			continue
		}

		if backtrack(0, equation.nums[0], equation.target, equation.nums) {
			result += equation.target
		}
	}

	return result
}
