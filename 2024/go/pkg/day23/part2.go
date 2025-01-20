package day23

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func findNextGroups(connections map[string]map[string]bool, groups map[string][]string, names map[string]struct{}) map[string][]string {
	nextGroups := make(map[string][]string)

	for _, group := range groups {
	potentialConnections:
		for potentialConnectionName := range names {
			if slices.Contains(group, potentialConnectionName) {
				// already in the group
				continue
			}

			for _, groupConnection := range group {
				if !connections[potentialConnectionName][groupConnection] {
					continue potentialConnections
				}
			}

			newGroup := make([]string, len(group)+1)
			copy(newGroup, group)
			newGroup[len(group)] = potentialConnectionName
			sort.Strings(newGroup)
			newGroupKey := strings.Join(newGroup, ",")
			nextGroups[newGroupKey] = newGroup
		}
	}

	return nextGroups
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	connections := make(map[string]map[string]bool)
	connectionNames := make(map[string]struct{})

	for i := range input {
		r := regexp.MustCompile("[a-z]+")
		names := r.FindAllString(input[i], -1)

		if _, ok := connections[names[0]]; !ok {
			connections[names[0]] = make(map[string]bool)
		}
		connections[names[0]][names[1]] = true

		if _, ok := connections[names[1]]; !ok {
			connections[names[1]] = make(map[string]bool)
		}
		connections[names[1]][names[0]] = true

		connectionNames[names[0]] = struct{}{}
		connectionNames[names[1]] = struct{}{}
	}

	groups := make(map[string][]string)

	fmt.Println("finding groups of length 1")
	for connectionName := range connectionNames {
		group := make([]string, 0)
		group = append(group, connectionName)
		groups[connectionName] = group
	}
	fmt.Printf("groups of length 1: %d\n", len(groups))

	length := 2
	for len(groups) != 1 {
		fmt.Printf("finding groups of length %d\n", length)
		groups = findNextGroups(connections, groups, connectionNames)
		fmt.Printf("groups of length %d: %d\n", length, len(groups))
		length += 1
	}

	for group := range groups {
		fmt.Println(group)
	}

	return 0
}
