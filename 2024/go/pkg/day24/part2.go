package day24

import (
	"fmt"
	"github.com/timredband/advent-of-code/pkg/utils"
	"math"
	"os"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	wiresByName, instructions := initialize(input)

	processInstructions(wiresByName, instructions)

	i := 0
	zValue := 0

	for {
		key := fmt.Sprintf("z%02d", i)

		if value, ok := wiresByName[key]; ok {
			zValue += (value * int(math.Pow(2, float64(i))))
			i += 1
		} else {
			break
		}
	}

	// x and y values
	i = 0
	xValue := 0
	yValue := 0

	for {
		xKey := fmt.Sprintf("x%02d", i)
		yKey := fmt.Sprintf("y%02d", i)

		if value, ok := wiresByName[xKey]; ok {
			xValue += (value * int(math.Pow(2, float64(i))))
		}

		if value, ok := wiresByName[yKey]; ok {
			yValue += (value * int(math.Pow(2, float64(i))))
		} else {
			break
		}

		i += 1
	}

	fmt.Printf("xValue: %046b, %d\n", xValue, xValue)
	fmt.Printf("yValue: %046b, %d\n", yValue, yValue)
	fmt.Printf("+Value: %046b, %d\n", xValue+yValue, xValue+yValue)
	fmt.Printf("zValue: %046b, %d\n", zValue, zValue)

	return 0
}
