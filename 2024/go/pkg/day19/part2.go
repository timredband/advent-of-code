package day19

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	towels := make(map[string]struct{})
	for _, v := range strings.Split(input[0], ", ") {
		towels[v] = struct{}{}
	}

	input = input[1:]

	cache := make([]int, 0)

	var DFS func(towel string, index int) int
	DFS = func(towel string, index int) int {
		if cache[index] != -1 {
			return cache[index]
		}

		count := 0

		for i := index; i < len(towel); i += 1 {
			sub := towel[index : i+1]
			if _, ok := towels[sub]; ok {
				count += DFS(towel, i+1)
			}
		}

		cache[index] = count

		return count
	}

	result := 0

	for i := range input {
		cache = make([]int, len(input[i])+1)
		for j := range len(cache) {
			cache[j] = -1
		}
		cache[len(cache)-1] = 1
		count := DFS(input[i], 0)
		result += count
	}

	return result
}
