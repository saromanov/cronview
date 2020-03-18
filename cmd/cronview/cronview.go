package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/saromanov/cronview/pkg/cmd"
)

func main() {
	if err := cmd.Build(os.Args); err != nil {
		logrus.Fatalf("unable to run project project: %v", err)
	}
}