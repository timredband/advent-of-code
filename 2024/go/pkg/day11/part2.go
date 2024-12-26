package day11

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type cacheEntry struct {
	blinks int
	stone  string
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)[0]
	stones := strings.Split(input, " ")
	cache := make(map[cacheEntry]int)
	blinks := 75

	var DFS func(stone string, blinks int) int
	DFS = func(stone string, blinks int) int {
		cacheEntry := cacheEntry{blinks: blinks, stone: stone}

		if stones, ok := cache[cacheEntry]; ok {
			return stones
		}

		if blinks == 0 {
			cache[cacheEntry] = 1
			return cache[cacheEntry]
		}

		if stones, ok := cache[cacheEntry]; ok {
			return stones
		}

		if stone == "0" {
			result := DFS("1", blinks-1)
			cache[cacheEntry] = result
			return cache[cacheEntry]
		}

		if len(stone)%2 == 0 {
			left := stone[0 : len(stone)/2]
			r, _ := strconv.Atoi(stone[len(stone)/2:])
			right := strconv.Itoa(r)

			leftResult := DFS(left, blinks-1)
			rightResult := DFS(right, blinks-1)

			cache[cacheEntry] = leftResult + rightResult
			return cache[cacheEntry]
		}

		value, _ := strconv.Atoi(stone)
		value *= 2024

		result := DFS(strconv.Itoa(value), blinks-1)
		cache[cacheEntry] = result

		return cache[cacheEntry]
	}

	result := 0

	for _, stone := range stones {
		result += DFS(stone, blinks)
	}

	return result
}
