package day13

import (
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/timredband/advent-of-code/pkg/utils"
)

// I had to lookup the idea for this one. Credit here https://advent-of-code.xavd.id/writeups/2024/day/13/.
func Part2(file *os.File) int {
	input := utils.ReadFile(file)

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
			Prize: Value{X: prizeX + 10000000000000, Y: prizeY + 10000000000000},
		}

		convertedaX := machine.A.X * machine.B.Y
		convertedXPrize := machine.Prize.X * machine.B.Y

		convertedaY := machine.A.Y * machine.B.X
		convertedYPrize := machine.Prize.Y * machine.B.X

		convertedA := convertedaX - convertedaY
		convertedPrize := convertedXPrize - convertedYPrize

		aPresses := float64(convertedPrize) / float64(convertedA)
		bPresses := (float64(machine.Prize.X) - (aPresses * float64(machine.A.X))) / float64(machine.B.X)

		if aPresses != math.Floor(aPresses) || bPresses != math.Floor(bPresses) {
			continue
		}

		tokens := (aPresses * 3) + (bPresses)
		result += int(tokens)
	}

	return result
}
