package day20

import (
	"math"
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

type entry struct {
	start coordinate
	end   coordinate
}

func Part2(file *os.File) int {
	input := utils.ReadFile(file)
	board := make([][]int, 0, len(input))

	start := coordinate{}
	end := coordinate{}

	for i := range input {
		row := make([]int, len(input[i]))

		for j := range row {
			if string(input[i][j]) == "#" {
				row[j] = 100
				row[j] = math.MaxInt
			}

			if string(input[i][j]) == "S" {
				start.x = i
				start.y = j
			}

			if string(input[i][j]) == "E" {
				end.x = i
				end.y = j
			}
		}

		board = append(board, row)
	}

	position := start

	for position != end {
		// north
		if board[position.x-1][position.y] == 0 && (coordinate{x: position.x - 1, y: position.y} != start) {
			board[position.x-1][position.y] = board[position.x][position.y] + 1
			position.x -= 1
		}

		// south
		if board[position.x+1][position.y] == 0 && (coordinate{x: position.x + 1, y: position.y} != start) {
			board[position.x+1][position.y] = board[position.x][position.y] + 1
			position.x += 1
		}

		// east
		if board[position.x][position.y+1] == 0 && (coordinate{x: position.x, y: position.y + 1} != start) {
			board[position.x][position.y+1] = board[position.x][position.y] + 1
			position.y += 1
		}

		// west
		if board[position.x][position.y-1] == 0 && (coordinate{x: position.x, y: position.y - 1} != start) {
			board[position.x][position.y-1] = board[position.x][position.y] + 1
			position.y -= 1
		}
	}

	cheats := make(map[entry]struct{})

	var DFS func(visited [][]int, start coordinate, steps int, c coordinate)
	DFS = func(visited [][]int, start coordinate, steps int, current coordinate) {
		if steps == 20 {
			return
		}

		startValue := board[start.x][start.y]
		value := startValue + steps

		if value > visited[current.x][current.y] {
			return
		}

		visited[current.x][current.y] = value

		// north cheating
		if current.x-1 > -1 && value+1 < visited[current.x-1][current.y] {
			next := current
			next.x -= 1
			DFS(visited, start, steps+1, next)
		}

		// south cheating
		if current.x+1 < len(board) && value+1 < visited[current.x+1][current.y] {
			next := current
			next.x += 1
			DFS(visited, start, steps+1, next)
		}

		// east cheating
		if current.y+1 < len(board[current.x]) && value+1 < visited[current.x][current.y+1] {
			next := current
			next.y += 1
			DFS(visited, start, steps+1, next)
		}

		// west cheating
		if current.y-1 > -1 && value+1 < visited[current.x][current.y-1] {
			next := current
			next.y -= 1
			DFS(visited, start, steps+1, next)
		}

		pico := 100

		// north exit
		if current.x-1 > -1 && board[current.x-1][current.y] != math.MaxInt && value+1 < board[current.x-1][current.y] {
			diff := board[current.x-1][current.y] - value - 1
			if diff >= pico {
				cheats[entry{start: start, end: coordinate{x: current.x - 1, y: current.y}}] = struct{}{}
			}
		}

		// south exit
		if current.x+1 < len(board) && board[current.x+1][current.y] != math.MaxInt && value+1 < board[current.x+1][current.y] {
			diff := board[current.x+1][current.y] - value - 1
			if diff >= pico {
				cheats[entry{start: start, end: coordinate{x: current.x + 1, y: current.y}}] = struct{}{}
			}
		}

		// east exit
		if current.y+1 < len(board[current.x]) && board[current.x][current.y+1] != math.MaxInt && value+1 < board[current.x][current.y+1] {
			diff := board[current.x][current.y+1] - value - 1
			if diff >= pico {
				cheats[entry{start: start, end: coordinate{x: current.x, y: current.y + 1}}] = struct{}{}
			}
		}

		// west exit
		if current.y-1 > -1 && board[current.x][current.y-1] != math.MaxInt && value+1 < board[current.x][current.y-1] {
			diff := board[current.x][current.y-1] - value - 1
			if diff >= pico {
				cheats[entry{start: start, end: coordinate{x: current.x, y: current.y - 1}}] = struct{}{}
			}
		}
	}

	position = start

	for position != end {
		next := coordinate{}

		// north
		if board[position.x-1][position.y] == board[position.x][position.y]+1 {
			next.x = position.x - 1
			next.y = position.y
		}

		// south
		if board[position.x+1][position.y] == board[position.x][position.y]+1 {
			next.x = position.x + 1
			next.y = position.y
		}

		// east
		if board[position.x][position.y+1] == board[position.x][position.y]+1 {
			next.x = position.x
			next.y = position.y + 1
		}

		// west
		if board[position.x][position.y-1] == board[position.x][position.y]+1 {
			next.x = position.x
			next.y = position.y - 1
		}

		v := make([][]int, 0, len(board))
		for i := range board {
			row := make([]int, len(board[i]))
			for j := range row {
				row[j] = math.MaxInt
			}
			v = append(v, row)
		}

		DFS(v, position, 0, position)

		position = next
	}

	// for i := range cheats {
	// 	fmt.Printf("s: %d,%d e: %d,%d\n", i.start.x, i.start.y, i.end.x, i.end.y)
	// }

	return len(cheats)
}
