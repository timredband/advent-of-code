package day17

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// this is not a solution. TODO: figure this out
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	r := regexp.MustCompile("[0-9]+")

	program := make([]int, 0)

	nums := r.FindAllString(input[3], -1)
	for i := range nums {
		num, _ := strconv.Atoi(nums[i])
		program = append(program, num)
	}

	output := ""
	result := 0

	expected := strings.Join(nums, "")
	reversed := make([]string, 0)

	for i := range expected {
		reversed = append(reversed, string(expected[len(expected)-i-1]))
	}

	reversedExpected := strings.Join(reversed, "")

	fmt.Println(reversedExpected)

	var calculate func(depth int, input int) bool
	calculate = func(depth int, input int) bool {
		registerA := input
		registerB := 0
		registerC := 0
		outputs := make([]string, 0)

		if registerA == 0 {
			outputs = append(outputs, strconv.Itoa(0))
		} else {
			for registerA != 0 {
				registerB = registerA % 8
				registerB = registerB ^ 1
				registerC = registerA / (int(math.Pow(2, float64(registerB))))
				registerB = registerB ^ registerC
				registerB = registerB ^ 4
				registerA = registerA >> 3
				outputs = append(outputs, strconv.Itoa(registerB%8))
			}
		}

		output = strings.Join(outputs, "")
		fmt.Println(output)

		return output == expected[depth:]
	}

	var DFS func(depth int, value int) bool
	DFS = func(depth int, value int) bool {
		if depth == 0 {
			return true
		}

		if output == expected {
			result = min(result, value)
			return true
		}

		for i := range 8 {
			v := ((i | 4) << (3 * depth)) | value
			if i < 4 {
				v = v ^ (4 << (3 * depth))
			}
			matches := calculate(depth, v)
			if matches {
				DFS(depth-1, v)
			}

			if output == reversedExpected {
				result = min(result, value)
				return true
			}
		}

		return false
	}

	// start := 0b100_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000
	start := 0
	DFS(15, start)

	fmt.Println("o: " + output)
	fmt.Println("e: " + expected)

	return result
}
