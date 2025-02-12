package day07

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Equation struct {
	target int
	nums   []int
}

func Part1(file *os.File) int {
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

	cache := make(map[string]struct{})

	var backtrack func(index int, value int, target int, nums []int) bool
	backtrack = func(index int, value int, target int, nums []int) bool {
		if index == len(nums)-1 && value == target {
			return true
		}

		index_converted := strconv.Itoa(index)
		value_converted := strconv.Itoa(value)

		key := index_converted + ":" + value_converted

		if _, ok := cache[key]; ok {
			return false
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

		cache[key] = struct{}{}

		return false
	}

	result := 0

	for _, equation := range equations {
		clear(cache)

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
