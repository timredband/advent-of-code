package day14

import (
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type position struct {
	x int
	y int
}

func Part1(file *os.File) int {
	input := utils.ReadFile(file)

	positions := make([]position, 0, len(input))
	xMax := 101
	yMax := 103

	for i := range input {
		r := regexp.MustCompile("-?[0-9]+")
		nums := r.FindAllString(input[i], -1)

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		xVelocity, _ := strconv.Atoi(nums[2])
		yVelocity, _ := strconv.Atoi(nums[3])

		for j := 0; j < 100; j += 1 {
			x = x + xVelocity
			if x < 0 {
				x *= -1
				x = x % xMax
				x = xMax - x
			} else {
				x = x % xMax
			}

			y = y + yVelocity
			if y < 0 {
				y *= -1
				y = y % yMax
				y = yMax - y
			} else {
				y = y % yMax
			}
		}

		positions = append(positions, position{x: x, y: y})
	}

	xMid := (xMax - 1) / 2
	yMid := (yMax - 1) / 2

	topLeftCount := 0
	topRightCount := 0
	bottomLeftCount := 0
	bottomRightCount := 0

	for _, position := range positions {
		if position.x == xMid || position.y == yMid {
			continue
		}

		if position.x < xMid && position.y < yMid {
			topLeftCount += 1
			continue
		}

		if position.x > xMid && position.y < yMid {
			topRightCount += 1
			continue
		}

		if position.x < xMid && position.y > yMid {
			bottomLeftCount += 1
			continue
		}

		if position.x > xMid && position.y > yMid {
			bottomRightCount += 1
			continue
		}
	}

	result := topLeftCount * topRightCount * bottomLeftCount * bottomRightCount
	return result
}
