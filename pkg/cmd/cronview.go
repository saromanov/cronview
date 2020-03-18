package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/saromanov/cronview/pkg/files"
	"github.com/urfave/cli/v2"
)

// Build provides handling of the commands
func Build(args []string) error {
	app := &cli.App{
		Name:  "cronview",
		Usage: "create puppet for the project",
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "add provides adding of the new command",
				Action: add,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func add(c *cli.Context) error {
	line := c.String("line")
	if line == "" {
		return fmt.Errorf("line is not defined")
	}
	mod, err := prepareAdd(line)
	if err != nil {
		return err
	}

	if err := files.Write(mod); err != nil {
		return fmt.Errorf("unable to write to crontab file")
	}
	return nil
}

func prepareAdd(line string) ([]string, error) {
	splitter := strings.Split(line, "  ")
	if len(splitter) == 0 {
		return nil, fmt.Errorf("unable to parse input data")
	}
	return []string{line}, nil
}
