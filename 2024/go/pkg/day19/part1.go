package day19

import (
	"os"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	towels := make(map[string]struct{})
	for _, v := range strings.Split(input[0], ", ") {
		towels[v] = struct{}{}
	}

	input = input[1:]

	var DFS func(towel string, index int) bool
	DFS = func(towel string, index int) bool {
		if index == len(towel) {
			return true
		}

		for i := len(towel); i > index; i -= 1 {
			if _, ok := towels[towel[index:i]]; ok {
				match := DFS(towel, i)
				if match {
					return true
				}
			}
		}

		return false
	}

	result := 0

	for i := range input {
		match := DFS(input[i], 0)

		if match {
			result += 1
		}
	}

	return result
}
