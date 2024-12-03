package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/timredband/advent-of-code/pkg/day1"
	"github.com/urfave/cli/v3"
)

func main() {
	var day string
	var part string
	var file string

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "d",
				Destination: &day,
			},
			&cli.StringFlag{
				Name:        "p",
				Destination: &part,
			},
			&cli.StringFlag{
				Name:        "f",
				Destination: &file,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			switch day {
			case "1":
				if file == "" {
					file = fmt.Sprintf("./inputs/day%s/part%s.txt", day, part)
				}

				fmt.Println(file)

				file, err := os.Open(file)
				if err != nil {
					return err
				}

				defer file.Close()

				day1.Execute(part, file)
			default:
				msg := fmt.Sprintf("unknown day: %s", day)
				return errors.New(msg)
			}

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
