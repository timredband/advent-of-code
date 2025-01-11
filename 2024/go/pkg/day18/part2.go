package day18

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	// size := 7
	size := 71
	end := 1
	blockingCoordinate := coordinate{}

	for {
		board := make([][]string, 0, size)
		visited := make([][]int, 0, len(board))

		for range size {
			row := make([]string, size)
			visitedRow := make([]int, size)
			for j := range size {
				row[j] = "."
				visitedRow[j] = math.MaxInt
			}
			board = append(board, row)
			visited = append(visited, visitedRow)
		}

		r := regexp.MustCompile("[0-9]+")

		for i := range input {
			if i == end {
				break
			}

			nums := r.FindAllString(input[i], -1)
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])

			board[y][x] = "#"
			blockingCoordinate.x = x
			blockingCoordinate.y = y
		}

		visited[size-1][size-1] = 0

		queue := make([]coordinate, 0)
		queue = append(queue, coordinate{x: size - 1, y: size - 1 - 1, value: 0}, coordinate{x: size - 1 - 1, y: size - 1, value: 0})

		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			if current.y < 0 || current.y > size-1 || current.x < 0 || current.x > size-1 {
				continue
			}

			if board[current.y][current.x] == "#" {
				continue
			}

			minimum := math.MaxInt

			// north
			if current.y-1 > -1 {
				minimum = min(minimum, visited[current.y-1][current.x])
			}

			// south
			if current.y+1 < size {
				minimum = min(minimum, visited[current.y+1][current.x])
			}

			// east
			if current.x+1 < size {
				minimum = min(minimum, visited[current.y][current.x+1])
			}

			// west
			if current.x-1 > -1 {
				minimum = min(minimum, visited[current.y][current.x-1])
			}

			if 1+minimum < visited[current.y][current.x] {
				visited[current.y][current.x] = 1 + minimum
				queue = append(queue,
					coordinate{y: current.y - 1, x: current.x},
					coordinate{y: current.y + 1, x: current.x},
					coordinate{y: current.y, x: current.x + 1},
					coordinate{y: current.y, x: current.x - 1},
				)
			}
		}

		if visited[0][0] == math.MaxInt {
			break
		}

		end += 1
	}

	fmt.Printf("%d,%d\n", blockingCoordinate.x, blockingCoordinate.y)

	return 0
}
