package cronview

import (
	"os"
	"github.com/saromanov/cronview/pkg/cmd"
)

func main() {
	if err := cmd.Build(os.Args); err != nil {
		logrus.Fatalf("unable to run project project: %v", err)
	}
}