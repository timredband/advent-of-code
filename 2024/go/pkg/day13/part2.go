package day13

import (
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

func Part2(file *os.File) int {
	input := utils.ReadFile(file)

	minimumTokens := math.MaxInt

	cache := make(map[CacheEntry]struct{})

	var backtrack func(machine Machine, xTotal int, yTotal, aPresses int, bPresses int)
	backtrack = func(machine Machine, xTotal int, yTotal, aPresses int, bPresses int) {
		if _, ok := cache[CacheEntry{aPresses: aPresses, bPresses: bPresses}]; ok {
			return
		}

		cache[CacheEntry{aPresses: aPresses, bPresses: bPresses}] = struct{}{}

		if xTotal == machine.Prize.X && yTotal == machine.Prize.Y {
			tokens := (aPresses * 3) + (bPresses)
			minimumTokens = min(minimumTokens, tokens)

			return
		}

		if xTotal > machine.Prize.X {
			return
		}

		if yTotal > machine.Prize.Y {
			return
		}

		// Press A button
		backtrack(machine, xTotal+machine.A.X, yTotal+machine.A.Y, aPresses+1, bPresses)

		// Press B button
		backtrack(machine, xTotal+machine.B.X, yTotal+machine.B.Y, aPresses, bPresses+1)
	}

	result := 0

	for i := 0; i < len(input); i += 3 {
		r := regexp.MustCompile("[0-9]+")

		a := r.FindAllString(input[i], -1)
		aX, _ := strconv.Atoi(a[0])
		aY, _ := strconv.Atoi(a[1])

		b := r.FindAllString(input[i+1], -1)
		bX, _ := strconv.Atoi(b[0])
		bY, _ := strconv.Atoi(b[1])

		prize := r.FindAllString(input[i+2], -1)
		prizeX, _ := strconv.Atoi(prize[0])
		prizeY, _ := strconv.Atoi(prize[1])

		machine := Machine{
			A:     Value{X: aX, Y: aY},
			B:     Value{X: bX, Y: bY},
			Prize: Value{X: prizeX + 100000, Y: prizeY + 100000},
		}

		backtrack(machine, 0, 0, 0, 0)

		if minimumTokens != math.MaxInt {
			result += minimumTokens
		}

		minimumTokens = math.MaxInt
		clear(cache)
	}

	return result
}
