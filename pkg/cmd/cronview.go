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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "add",
				Value: "",
				Usage: "adding of the new scenario",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}
