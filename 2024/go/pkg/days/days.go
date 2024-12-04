package days

import (
	"errors"
	"fmt"
	"os"

	"github.com/timredband/advent-of-code/pkg/day1"
	"github.com/timredband/advent-of-code/pkg/day2"
)

func Execute(day string, part string, file *os.File) error {
	switch day {
	case "1":
		return day1.Execute(part, file)
	case "2":
		return day2.Execute(part, file)
	default:
		return errors.New(fmt.Sprintf("unknown day: %s", day))
	}
}