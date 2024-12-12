package days

import (
	"errors"
	"fmt"
	"os"

	"github.com/timredband/advent-of-code/pkg/day1"
	"github.com/timredband/advent-of-code/pkg/day2"
	"github.com/timredband/advent-of-code/pkg/day3"
	"github.com/timredband/advent-of-code/pkg/day4"
	"github.com/timredband/advent-of-code/pkg/day5"
	"github.com/timredband/advent-of-code/pkg/day6"
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
	default:
		return 0, errors.New(fmt.Sprintf("unknown day: %s", day))
	}
}
