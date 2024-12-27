package day12

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type Coordinate struct {
	x int
	y int
}

type Perimeter struct {
	up    bool
	down  bool
	left  bool
	right bool
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]string, len(input))
	visited := make([][]bool, len(input))

	for i := range input {
		board[i] = make([]string, len(input[i]))
		visited[i] = make([]bool, len(input[i]))

		for j := range input[i] {
			board[i][j] = string(input[i][j])
		}
	}

	var DFS func(i int, j int, current string, region map[Coordinate]Perimeter) (area int)
	DFS = func(i int, j int, current string, region map[Coordinate]Perimeter) (area int) {
		if visited[i][j] {
			return
		}

		visited[i][j] = true

		area += 1

		coordinate := Coordinate{x: i, y: j}
		perimeter := Perimeter{}

		// up
		if i-1 > -1 && board[i-1][j] == current {
			upArea := DFS(i-1, j, current, region)
			area += upArea
		} else {
			perimeter.up = true
		}

		// down
		if i+1 < len(board) && board[i+1][j] == current {
			downArea := DFS(i+1, j, current, region)
			area += downArea
		} else {
			perimeter.down = true
		}

		// left
		if j-1 > -1 && board[i][j-1] == current {
			leftArea := DFS(i, j-1, current, region)
			area += leftArea
		} else {
			perimeter.left = true
		}

		// right
		if j+1 < len(board[i]) && board[i][j+1] == current {
			rightArea := DFS(i, j+1, current, region)
			area += rightArea
		} else {
			perimeter.right = true
		}

		if perimeter.up || perimeter.down || perimeter.left || perimeter.right {
			region[coordinate] = perimeter
		}

		return
	}

	result := 0

	region := make(map[Coordinate]Perimeter)

	for i := range board {
		for j := range board[i] {
			area := DFS(i, j, board[i][j], region)

			sides := 0
			for k, v := range region {
				kOrig := k
				vOrig := v

				if v.up {
					// left
					for {
						v.up = false
						leftCoordinate := Coordinate{x: k.x, y: k.y - 1}
						if nextPerimeter, ok := region[leftCoordinate]; ok && nextPerimeter.up {
							region[k] = v
							k = leftCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig

					// right
					for {
						v.up = false
						rightCoordinate := Coordinate{x: k.x, y: k.y + 1}
						if nextPerimeter, ok := region[rightCoordinate]; ok && nextPerimeter.up {
							region[k] = v
							k = rightCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig
					sides += 1
				}

				if v.down {
					// left
					for {
						v.down = false
						leftCoordinate := Coordinate{x: k.x, y: k.y - 1}
						if nextPerimeter, ok := region[leftCoordinate]; ok && nextPerimeter.down {
							region[k] = v
							k = leftCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig

					// right
					for {
						v.down = false
						rightCoordinate := Coordinate{x: k.x, y: k.y + 1}
						if nextPerimeter, ok := region[rightCoordinate]; ok && nextPerimeter.down {
							region[k] = v
							k = rightCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig
					sides += 1
				}

				if v.left {
					// up
					for {
						v.left = false
						upCoordinate := Coordinate{x: k.x - 1, y: k.y}
						if nextPerimeter, ok := region[upCoordinate]; ok && nextPerimeter.left {
							region[k] = v
							k = upCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig

					// down
					for {
						v.left = false
						rightCoordinate := Coordinate{x: k.x + 1, y: k.y}
						if nextPerimeter, ok := region[rightCoordinate]; ok && nextPerimeter.left {
							region[k] = v
							k = rightCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig
					sides += 1
				}

				if v.right {
					// up
					for {
						v.right = false
						upCoordinate := Coordinate{x: k.x - 1, y: k.y}
						if nextPerimeter, ok := region[upCoordinate]; ok && nextPerimeter.right {
							region[k] = v
							k = upCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					v = vOrig
					k = kOrig

					// down
					for {
						v.right = false
						rightCoordinate := Coordinate{x: k.x + 1, y: k.y}
						if nextPerimeter, ok := region[rightCoordinate]; ok && nextPerimeter.right {
							region[k] = v
							k = rightCoordinate
							v = nextPerimeter
						} else {
							region[k] = v
							break
						}
					}

					sides += 1
				}
			}

			region = make(map[Coordinate]Perimeter)
			result += (area * sides)
		}
	}

	return result
}
