package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/timredband/advent-of-code/pkg/days"
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
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "p",
				Destination: &part,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "f",
				Destination: &file,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if file == "" {
				file = fmt.Sprintf("./inputs/day%s/input.txt", day)
			}

			file, err := os.Open(file)
			if err != nil {
				return err
			}

			defer file.Close()

			result, err := days.Execute(day, part, file)
			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
