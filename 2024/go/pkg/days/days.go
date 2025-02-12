package days

import (
	"errors"
	"fmt"
	"os"

	"github.com/timredband/advent-of-code/pkg/day01"
	"github.com/timredband/advent-of-code/pkg/day02"
	"github.com/timredband/advent-of-code/pkg/day03"
	"github.com/timredband/advent-of-code/pkg/day04"
	"github.com/timredband/advent-of-code/pkg/day05"
	"github.com/timredband/advent-of-code/pkg/day06"
	"github.com/timredband/advent-of-code/pkg/day07"
	"github.com/timredband/advent-of-code/pkg/day08"
	"github.com/timredband/advent-of-code/pkg/day09"
	"github.com/timredband/advent-of-code/pkg/day10"
	"github.com/timredband/advent-of-code/pkg/day11"
	"github.com/timredband/advent-of-code/pkg/day12"
	"github.com/timredband/advent-of-code/pkg/day13"
	"github.com/timredband/advent-of-code/pkg/day14"
	"github.com/timredband/advent-of-code/pkg/day15"
	"github.com/timredband/advent-of-code/pkg/day16"
	"github.com/timredband/advent-of-code/pkg/day17"
	"github.com/timredband/advent-of-code/pkg/day18"
	"github.com/timredband/advent-of-code/pkg/day19"
	"github.com/timredband/advent-of-code/pkg/day20"
	"github.com/timredband/advent-of-code/pkg/day21"
	"github.com/timredband/advent-of-code/pkg/day22"
	"github.com/timredband/advent-of-code/pkg/day23"
	"github.com/timredband/advent-of-code/pkg/day24"
	"github.com/timredband/advent-of-code/pkg/day25"
)

func CreateExecutor(part1 func(file *os.File) int, part2 func(file *os.File) int) func(part string, file *os.File) (int, error) {
	executor := func(part string, file *os.File) (int, error) {
		switch part {
		case "1":
			return part1(file), nil
		case "2":
			return part2(file), nil
		default:
			return 0, errors.New(fmt.Sprintf("unknown part: %s", part))
		}
	}

	return executor
}

func Execute(day string, part string, file *os.File) (int, error) {
	switch day {
	case "1":
		return CreateExecutor(day01.Part1, day01.Part2)(part, file)
	case "2":
		return CreateExecutor(day02.Part1, day02.Part2)(part, file)
	case "3":
		return CreateExecutor(day03.Part1, day03.Part2)(part, file)
	case "4":
		return CreateExecutor(day04.Part1, day04.Part2)(part, file)
	case "5":
		return CreateExecutor(day05.Part1, day05.Part2)(part, file)
	case "6":
		return CreateExecutor(day06.Part1, day06.Part2)(part, file)
	case "7":
		return CreateExecutor(day07.Part1, day07.Part2)(part, file)
	case "8":
		return CreateExecutor(day08.Part1, day08.Part2)(part, file)
	case "9":
		return CreateExecutor(day09.Part1, day09.Part2)(part, file)
	case "10":
		return CreateExecutor(day10.Part1, day10.Part2)(part, file)
	case "11":
		return CreateExecutor(day11.Part1, day11.Part2)(part, file)
	case "12":
		return CreateExecutor(day12.Part1, day12.Part2)(part, file)
	case "13":
		return CreateExecutor(day13.Part1, day13.Part2)(part, file)
	case "14":
		return CreateExecutor(day14.Part1, day14.Part2)(part, file)
	case "15":
		return CreateExecutor(day15.Part1, day15.Part2)(part, file)
	case "16":
		return CreateExecutor(day16.Part1, day16.Part2)(part, file)
	case "17":
		return CreateExecutor(day17.Part1, day17.Part2)(part, file)
	case "18":
		return CreateExecutor(day18.Part1, day18.Part2)(part, file)
	case "19":
		return CreateExecutor(day19.Part1, day19.Part2)(part, file)
	case "20":
		return CreateExecutor(day20.Part1, day20.Part2)(part, file)
	case "21":
		return CreateExecutor(day21.Part1, day21.Part2)(part, file)
	case "22":
		return CreateExecutor(day22.Part1, day22.Part2)(part, file)
	case "23":
		return CreateExecutor(day23.Part1, day23.Part2)(part, file)
	case "24":
		return CreateExecutor(day24.Part1, day24.Part2)(part, file)
	case "25":
		return CreateExecutor(day25.Part1, day25.Part2)(part, file)
	default:
		return 0, errors.New(fmt.Sprintf("unknown day: %s", day))
	}
}
