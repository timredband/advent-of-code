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
	var path string

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
				Destination: &path,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if path == "" {
				path = fmt.Sprintf("./inputs/day%02s/input.txt", day)
			}

			result, err := days.Execute(day, part, path)
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
