package cmd

import (
	"strings"
	"fmt"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/saromanov/cronview/pkg/models"
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
				Action: build,
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
	return nil
}

func prepareAdd(line string)(*models.Crontab, error) {
	splitter := strings.Split(line, "  ")
	if len(splitter) == 0 {
		return nil, fmt.Errorf("unable to parse input data")
	}
	return &models.Crontab{
		Schedule: splitter[0],
		Command: splitter[1],
	}, nil
}
