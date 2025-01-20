package days

import (
	"errors"
	"fmt"
	"os"

	"github.com/timredband/advent-of-code/pkg/day1"
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
	"github.com/timredband/advent-of-code/pkg/day2"
	"github.com/timredband/advent-of-code/pkg/day20"
	"github.com/timredband/advent-of-code/pkg/day21"
	"github.com/timredband/advent-of-code/pkg/day22"
	"github.com/timredband/advent-of-code/pkg/day23"
	"github.com/timredband/advent-of-code/pkg/day24"
	"github.com/timredband/advent-of-code/pkg/day3"
	"github.com/timredband/advent-of-code/pkg/day4"
	"github.com/timredband/advent-of-code/pkg/day5"
	"github.com/timredband/advent-of-code/pkg/day6"
	"github.com/timredband/advent-of-code/pkg/day7"
	"github.com/timredband/advent-of-code/pkg/day8"
	"github.com/timredband/advent-of-code/pkg/day9"
)

func Execute(day string, part string, file *os.File) (int, error) {
	switch day {
	case "1":
		return day1.Execute(part, file)
	case "2":
		return day2.Execute(part, file)
	case "3":
		return day3.Execute(part, file)
	case "4":
		return day4.Execute(part, file)
	case "5":
		return day5.Execute(part, file)
	case "6":
		return day6.Execute(part, file)
	case "7":
		return day7.Execute(part, file)
	case "8":
		return day8.Execute(part, file)
	case "9":
		return day9.Execute(part, file)
	case "10":
		return day10.Execute(part, file)
	case "11":
		return day11.Execute(part, file)
	case "12":
		return day12.Execute(part, file)
	case "13":
		return day13.Execute(part, file)
	case "14":
		return day14.Execute(part, file)
	case "15":
		return day15.Execute(part, file)
	case "16":
		return day16.Execute(part, file)
	case "17":
		return day17.Execute(part, file)
	case "18":
		return day18.Execute(part, file)
	case "19":
		return day19.Execute(part, file)
	case "20":
		return day20.Execute(part, file)
	case "21":
		return day21.Execute(part, file)
	case "22":
		return day22.Execute(part, file)
	case "23":
		return day23.Execute(part, file)
	case "24":
		return day24.Execute(part, file)
	default:
		return 0, errors.New(fmt.Sprintf("unknown day: %s", day))
	}
}
