package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)
	result := 0

	var sb strings.Builder

	for _, line := range inputs {
		sb.WriteString(line)
	}

	concatenated := sb.String()

	doRegex, _ := regexp.Compile("do\\(\\)")
	input := doRegex.Split(concatenated, -1)

	for i := 0; i < len(input); i += 1 {
		dontRegex, _ := regexp.Compile("don't\\(\\)")
		dontSplit := dontRegex.Split(input[i], -1)

		mulRegex, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
		mulSplit := mulRegex.FindAllString(dontSplit[0], -1)
		fmt.Println(mulSplit)
		numRegex, _ := regexp.Compile("[0-9]+")

		for i := 0; i < len(mulSplit); i += 1 {
			nums := numRegex.FindAllString(mulSplit[i], -1)
			first, _ := strconv.Atoi(nums[0])
			second, _ := strconv.Atoi(nums[1])
			result += (first * second)
		}
	}

	return result
}
