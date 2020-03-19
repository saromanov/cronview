package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/saromanov/cronview/pkg/commands"
	"github.com/saromanov/cronview/pkg/files"
	"github.com/saromanov/cronview/pkg/models"
	"github.com/urfave/cli/v2"
)

// Build provides handling of the commands
func Build(args []string) error {
	app := &cli.App{
		Name:  "cronview",
		Usage: "create puppet for the project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "schedule",
				Value: "",
				Usage: "cron schedule",
			},
			&cli.StringFlag{
				Name:  "command",
				Usage: "command to the cron schedule",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "add provides adding of the new command",
				Action: add,
			},
			{
				Name:   "list",
				Usage:  "showing list of cron tasks for current user",
				Action: list,
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
	schedule := c.String("schedule")
	if schedule == "" {
		return fmt.Errorf("schedule is not defined")
	}
	command := c.String("command")
	if command == "" {
		return fmt.Errorf("command is not defined")
	}

	if err := validateSchedule(schedule); err != nil {
		return fmt.Errorf("unable to validate schedule: %v", err)
	}

	if err := files.Write(&models.Crontab{
		Records: []models.Record{models.Record{
			Schedule: schedule,
			Command:  command,
		},
		},
	}); err != nil {
		return fmt.Errorf("unable to write to crontab file: %v", err)
	}
	return nil
}

// return list of cron tasks
func list(c *cli.Context) error {
	name := "lst.tmp"
	if err := commands.ExecAndWrite(name); err != nil {
		return fmt.Errorf("unable to execute command: %v", err)
	}

	data, err := files.Read(name)
	if err != nil {
		return fmt.Errorf("unable to read from %s: %v", name, err)
	}
	for _, d := range data.Records {
		fmt.Println("DATA: ", d)
	}
	return nil
}

// providing of the validating cron schedule
func validateSchedule(s string) error {
	splitter := strings.Split(s, " ")
	if len(splitter) != 5 {
		return fmt.Errorf("invalid cron format")
	}
	return nil
}
