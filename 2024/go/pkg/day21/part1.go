package day21

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func createNumberPad() map[string]map[string]string {
	numberPad := make(map[string]map[string]string)

	numberPad["0"] = make(map[string]string)
	numberPad["0"]["A"] = ">"
	numberPad["0"]["2"] = "^"

	numberPad["1"] = make(map[string]string)
	numberPad["1"]["2"] = ">"
	numberPad["1"]["4"] = "^"

	numberPad["2"] = make(map[string]string)
	numberPad["2"]["0"] = "v"
	numberPad["2"]["1"] = "<"
	numberPad["2"]["3"] = ">"
	numberPad["2"]["5"] = "^"

	numberPad["3"] = make(map[string]string)
	numberPad["3"]["A"] = "v"
	numberPad["3"]["2"] = "<"
	numberPad["3"]["6"] = "^"

	numberPad["4"] = make(map[string]string)
	numberPad["4"]["1"] = "v"
	numberPad["4"]["5"] = ">"
	numberPad["4"]["7"] = "^"

	numberPad["5"] = make(map[string]string)
	numberPad["5"]["2"] = "v"
	numberPad["5"]["4"] = "<"
	numberPad["5"]["6"] = ">"
	numberPad["5"]["8"] = "^"

	numberPad["6"] = make(map[string]string)
	numberPad["6"]["3"] = "v"
	numberPad["6"]["5"] = "<"
	numberPad["6"]["9"] = "^"

	numberPad["7"] = make(map[string]string)
	numberPad["7"]["4"] = "v"
	numberPad["7"]["8"] = ">"

	numberPad["8"] = make(map[string]string)
	numberPad["8"]["5"] = "v"
	numberPad["8"]["7"] = "<"
	numberPad["8"]["9"] = ">"

	numberPad["9"] = make(map[string]string)
	numberPad["9"]["6"] = "v"
	numberPad["9"]["8"] = "<"

	numberPad["A"] = make(map[string]string)
	numberPad["A"]["0"] = "<"
	numberPad["A"]["3"] = "^"

	return numberPad
}

func createDirectionalPad() map[string]map[string]string {
	directionalPad := make(map[string]map[string]string)

	directionalPad["^"] = make(map[string]string)
	directionalPad["^"]["A"] = ">"
	directionalPad["^"]["v"] = "v"

	directionalPad["v"] = make(map[string]string)
	directionalPad["v"]["<"] = "<"
	directionalPad["v"]["^"] = "^"
	directionalPad["v"][">"] = ">"

	directionalPad["<"] = make(map[string]string)
	directionalPad["<"]["v"] = ">"

	directionalPad[">"] = make(map[string]string)
	directionalPad[">"]["v"] = "<"
	directionalPad[">"]["A"] = "^"

	directionalPad["A"] = make(map[string]string)
	directionalPad["A"]["^"] = "<"
	directionalPad["A"][">"] = "v"

	return directionalPad
}

type search struct {
	paths []*strings.Builder
}

func newSearch() *search {
	s := search{}
	s.paths = make([]*strings.Builder, 0)
	return &s
}

func (s *search) DFS(pad map[string]map[string]string, current string, target string, sb *strings.Builder, visited []string) {
	if current == target {
		s.paths = append(s.paths, sb)
		return
	}

	if direction, ok := pad[current][target]; ok {
		sb.WriteString(direction)
		s.paths = append(s.paths, sb)
		return
	}

	for next, direction := range pad[current] {
		if slices.Contains(visited, next) {
			continue
		}
		visitedCopy := make([]string, len(visited)+1)
		copy(visitedCopy, visited)
		visitedCopy = append(visitedCopy, current)
		sbCopy := strings.Builder{}
		sbCopy.WriteString(sb.String() + direction)
		s.DFS(pad, next, target, &sbCopy, visitedCopy)
	}
}

func control(sequence string, pad map[string]map[string]string) []string {
	allBuilders := make([][]*strings.Builder, 0)
	current := "A"

	for _, target := range sequence {
		visited := make([]string, 0)

		s := newSearch()

		s.DFS(pad, current, string(target), &strings.Builder{}, visited)

		for i := range s.paths {
			s.paths[i].WriteString("A")
		}

		allBuilders = append(allBuilders, s.paths)

		current = string(target)
	}

	for len(allBuilders) != 1 {
		combined := make([]*strings.Builder, 0)
		for i := range allBuilders[0] {
			for j := range allBuilders[1] {
				sb := strings.Builder{}
				sb.WriteString(allBuilders[0][i].String())
				sb.WriteString(allBuilders[1][j].String())
				combined = append(combined, &sb)
			}
		}
		allBuilders[1] = combined
		allBuilders = allBuilders[1:]
	}

	allPaths := make([]string, 0)

	for i := range allBuilders[0] {
		allPaths = append(allPaths, allBuilders[0][i].String())
	}

	slices.SortFunc(allPaths, func(a string, b string) int {
		return len(a) - len(b)
	})

	minLength := len(allPaths[0])

	result := make([]string, 0)

	for i := range allPaths {
		if len(allPaths[i]) == minLength {
			result = append(result, allPaths[i])
		}
	}

	return result
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	numberPad := createNumberPad()
	directionalPad := createDirectionalPad()

	result := 0
	for _, code := range input {
		numericCode, _ := strconv.Atoi(code[0:3])
		minPath := ""
		numberPadPaths := control(code, numberPad)

		for _, numberPadPath := range numberPadPaths {
			directionalPadPathOnes := control(numberPadPath, directionalPad)

			for _, directionalPadPathOne := range directionalPadPathOnes {
				directionalPadPathTwos := control(directionalPadPathOne, directionalPad)
				for _, directionalPadPathTwo := range directionalPadPathTwos {
					if minPath == "" {
						minPath = directionalPadPathTwo
					}

					if len(directionalPadPathTwo) < len(minPath) {
						minPath = directionalPadPathTwo
					}
				}
			}
		}

		result += (numericCode * len(minPath))
	}

	return result
}
