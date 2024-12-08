package day5

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	inputs := utils.ReadFile(file)

	rules := make(map[string]struct{})
	rulesIndex := 0

	for i, v := range inputs {
		if !strings.Contains(v, "|") {
			rulesIndex = i
			break
		}

		rules[v] = struct{}{}
	}

	result := 0

	for i := rulesIndex; i < len(inputs); i += 1 {
		current := strings.Split(inputs[i], ",")
		middle := current[len(current)/2]
		valid := true

	outer:
		for j := 0; j < len(current); j += 1 {
			for k := j + 1; k < len(current); k += 1 {
				rule := string(current[j]) + "|" + string(current[k])
				if _, ok := rules[rule]; !ok {
					valid = false
					break outer
				}
			}
		}

		if valid {
			value, _ := strconv.Atoi(middle)
			result += value
		}
	}

	return result
}
