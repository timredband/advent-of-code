package day05

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
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
		valid := true

	outer:
		for j := 0; j < len(current); j += 1 {
			for k := j + 1; k < len(current); k += 1 {
				rule := string(current[j]) + "|" + string(current[k])
				if _, ok := rules[rule]; !ok {
					valid = false
					current[j], current[k] = current[k], current[j]
					goto outer // I feel dirty using a goto, but this doesn't seem that bad.
				}
			}
		}

		if !valid {
			middle := current[len(current)/2]
			value, _ := strconv.Atoi(middle)
			result += value
		}
	}

	return result
}
