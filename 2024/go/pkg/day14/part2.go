package day14

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type robot struct {
	x         int
	y         int
	xVelocity int
	yVelocity int
}

// this solution relies on that for the tree every robot is in a unique position. This seems lucky this is the case.
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	robots := make([]robot, 0, len(input))

	xMax := 101
	yMax := 103

	for i := range input {
		r := regexp.MustCompile("-?[0-9]+")
		nums := r.FindAllString(input[i], -1)

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		xVelocity, _ := strconv.Atoi(nums[2])
		yVelocity, _ := strconv.Atoi(nums[3])

		robots = append(robots, robot{x: x, y: y, xVelocity: xVelocity, yVelocity: yVelocity})
	}

	board := make([][]string, 0, yMax)
	for i := 0; i < yMax; i += 1 {
		row := make([]string, xMax)
		for j := range row {
			row[j] = " "

		}
		board = append(board, row)
	}

	for i := 0; i < 10000; i += 1 {
		for k := 0; k < yMax; k += 1 {
			board[k] = make([]string, xMax)
			for l := range board[k] {
				board[k][l] = " "
			}
		}

		uniquePositions := make(map[position]struct{})

		for j, robot := range robots {
			robot.x = robot.x + robot.xVelocity
			if robot.x < 0 {
				robot.x *= -1
				robot.x = robot.x % xMax
				robot.x = xMax - robot.x
			} else {
				robot.x = robot.x % xMax
			}

			robot.y = robot.y + robot.yVelocity
			if robot.y < 0 {
				robot.y *= -1
				robot.y = robot.y % yMax
				robot.y = yMax - robot.y
			} else {
				robot.y = robot.y % yMax
			}

			robots[j] = robot

			uniquePositions[position{x: robot.x, y: robot.y}] = struct{}{}
			board[robot.y][robot.x] = "*"
		}

		if len(uniquePositions) == len(robots) {
			for p := range board {
				fmt.Println(board[p])
			}
			return i + 1
		}
	}

	return 0
}
