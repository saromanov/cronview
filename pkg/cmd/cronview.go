package cmd

import (
	"os"
	"github.com/sirupsen/logrus"
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
	return nil
}
