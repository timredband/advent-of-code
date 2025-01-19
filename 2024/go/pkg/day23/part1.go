package day23

import (
	"os"
	"regexp"
	"slices"
	"sort"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	connectionsByName := make(map[string][]string)

	for i := range input {
		r := regexp.MustCompile("[a-z]+")
		names := r.FindAllString(input[i], -1)

		if _, ok := connectionsByName[names[0]]; !ok {
			connectionsByName[names[0]] = make([]string, 0, 1)
		}
		connectionsByName[names[0]] = append(connectionsByName[names[0]], names[1])

		if _, ok := connectionsByName[names[1]]; !ok {
			connectionsByName[names[1]] = make([]string, 0, 1)
		}
		connectionsByName[names[1]] = append(connectionsByName[names[1]], names[0])
	}

	c := make(map[string]struct{})

	for first, connections := range connectionsByName {
		for _, second := range connections {
			for _, third := range connectionsByName[second] {
				if slices.Contains(connections, third) {
					connected := []string{first, second, third}
					sort.Strings(connected)

					if string(connected[0][0]) == "t" || string(connected[1][0]) == "t" || string(connected[2][0]) == "t" {
						c[connected[0]+connected[1]+connected[2]] = struct{}{}
					}
				}
			}
		}
	}

	return len(c)
}
