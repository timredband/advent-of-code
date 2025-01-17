package day21

import (
	"os"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// couldn't figure out part 2. Used this post as a reference to solve it. https://www.reddit.com/r/adventofcode/comments/1hjx0x4/2024_day_21_quick_tutorial_to_solve_part_2_in/
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	result := solve(input, 25)

	return result
}
