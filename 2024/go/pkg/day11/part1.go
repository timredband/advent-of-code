package day11

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	input := utils.ReadFile(file)[0]
	stones := strings.Split(input, " ")

	blinks := 25

	for blink := 0; blink < blinks; blink += 1 {
		for i, stone := range stones {
			if stone == "0" {
				stones[i] = "1"
				continue
			}

			if len(stone)%2 == 0 {
				left := stone[0 : len(stone)/2]
				r, _ := strconv.Atoi(stone[len(stone)/2:])
				right := strconv.Itoa(r)
				stones[i] = left
				stones = append(stones, right)
				continue
			}

			value, _ := strconv.Atoi(stone)
			value *= 2024
			stones[i] = strconv.Itoa(value)
		}
	}

	result := len(stones)
	return result
}
