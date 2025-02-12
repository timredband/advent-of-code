package day06

import (
	"os"
	"strconv"
	"strings"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func createStartingBoard(initialBoard [][]string) [][]string {
	board := make([][]string, len(initialBoard))

	for i := range initialBoard {
		board[i] = make([]string, len(initialBoard[i]))
		copy(board[i], initialBoard[i])
	}

	return board
}

func walk(board [][]string, position Position) int {
	for position.i > 0 && position.j < len(board) {
		current := board[position.i][position.j]

		if string(current[0]) == "X" {
			//been here before

			number, _ := strconv.Atoi(string(current[1]))

			// hacky way to figure out if I've visited this position "many" times which might be a good indicator I'm in a loop. Obviously there's flaw with this.
			if number == 9 {
				return 1
			}

			number += 1
			board[position.i][position.j] = "X" + strconv.Itoa(number)
		} else {
			board[position.i][position.j] = "X1"
		}

		if position.direction == "north" {
			if position.i-1 < 0 {
				break
			}

			if board[position.i-1][position.j] == "#" {
				position.direction = "east"
			} else {
				position.i -= 1
			}

			continue
		}

		if position.direction == "south" {
			if position.i+1 > len(board)-1 {
				break
			}

			if board[position.i+1][position.j] == "#" {
				position.direction = "west"
			} else {
				position.i += 1
			}

			continue
		}

		if position.direction == "east" {
			if position.j+1 > len(board)-1 {
				break
			}

			if board[position.i][position.j+1] == "#" {
				position.direction = "south"
			} else {
				position.j += 1
			}

			continue

		}

		if position.direction == "west" {
			if position.j-1 < 0 {
				break
			}

			if board[position.i][position.j-1] == "#" {
				position.direction = "north"
			} else {
				position.j -= 1
			}

			continue
		}
	}

	return 0
}

func Part2(file *os.File) int {
	inputs := utils.ReadFile(file)
	initialBoard := make([][]string, len(inputs))
	initialPosition := Position{direction: "north"}

	for i, v := range inputs {
		initialBoard[i] = strings.Split(v, "")

		for j, c := range initialBoard[i] {
			if c == "^" {
				initialPosition.i = i
				initialPosition.j = j
			}
		}
	}

	result := 0

	for i := range initialBoard {
		for j := range initialBoard[i] {
			if initialBoard[i][j] == "#" {
				continue
			}

			board := createStartingBoard(initialBoard)
			board[i][j] = "#"
			position := initialPosition

			result += walk(board, position)
		}
	}

	return result
}
