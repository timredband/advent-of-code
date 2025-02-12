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

func CreateExecutor(part1 func(file *os.File) int, part2 func(file *os.File) int) func(part string, path string) (int, error) {
	executor := func(part string, path string) (int, error) {
		file, err := os.Open(path)

		if err != nil {
			return 0, err
		}

		defer file.Close()

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

func Execute(day string, part string, path string) (int, error) {
	switch day {
	case "1":
		executor := CreateExecutor(day01.Part1, day01.Part2)
		return executor(part, path)
	case "2":
		executor := CreateExecutor(day02.Part1, day02.Part2)
		return executor(part, path)
	case "3":
		executor := CreateExecutor(day03.Part1, day03.Part2)
		return executor(part, path)
	case "4":
		executor := CreateExecutor(day04.Part1, day04.Part2)
		return executor(part, path)
	case "5":
		executor := CreateExecutor(day05.Part1, day05.Part2)
		return executor(part, path)
	case "6":
		executor := CreateExecutor(day06.Part1, day06.Part2)
		return executor(part, path)
	case "7":
		executor := CreateExecutor(day07.Part1, day07.Part2)
		return executor(part, path)
	case "8":
		executor := CreateExecutor(day08.Part1, day08.Part2)
		return executor(part, path)
	case "9":
		executor := CreateExecutor(day09.Part1, day09.Part2)
		return executor(part, path)
	case "10":
		executor := CreateExecutor(day10.Part1, day10.Part2)
		return executor(part, path)
	case "11":
		executor := CreateExecutor(day11.Part1, day11.Part2)
		return executor(part, path)
	case "12":
		executor := CreateExecutor(day12.Part1, day12.Part2)
		return executor(part, path)
	case "13":
		executor := CreateExecutor(day13.Part1, day13.Part2)
		return executor(part, path)
	case "14":
		executor := CreateExecutor(day14.Part1, day14.Part2)
		return executor(part, path)
	case "15":
		executor := CreateExecutor(day15.Part1, day15.Part2)
		return executor(part, path)
	case "16":
		executor := CreateExecutor(day16.Part1, day16.Part2)
		return executor(part, path)
	case "17":
		executor := CreateExecutor(day17.Part1, day17.Part2)
		return executor(part, path)
	case "18":
		executor := CreateExecutor(day18.Part1, day18.Part2)
		return executor(part, path)
	case "19":
		executor := CreateExecutor(day19.Part1, day19.Part2)
		return executor(part, path)
	case "20":
		executor := CreateExecutor(day20.Part1, day20.Part2)
		return executor(part, path)
	case "21":
		executor := CreateExecutor(day21.Part1, day21.Part2)
		return executor(part, path)
	case "22":
		executor := CreateExecutor(day22.Part1, day22.Part2)
		return executor(part, path)
	case "23":
		executor := CreateExecutor(day23.Part1, day23.Part2)
		return executor(part, path)
	case "24":
		executor := CreateExecutor(day24.Part1, day24.Part2)
		return executor(part, path)
	case "25":
		executor := CreateExecutor(day25.Part1, day25.Part2)
		return executor(part, path)
	default:
		return 0, errors.New(fmt.Sprintf("unknown day: %s", day))
	}
}
