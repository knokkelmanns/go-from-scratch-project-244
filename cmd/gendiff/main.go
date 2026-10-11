// Command gendiff compares two configuration files and shows a difference.
package main

import (
	"context"
	"fmt"
	"os"

	"code"
	"code/formatters"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "output format",
				Value:   formatters.DefaultFormat,
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			if cmd.NArg() != 2 {
				return cli.ShowAppHelp(cmd)
			}

			result, err := code.GenDiff(cmd.Args().Get(0), cmd.Args().Get(1), cmd.String("format"))
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
