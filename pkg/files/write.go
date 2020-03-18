package files

import (
	"fmt"
	"os"

	"github.com/saromanov/cronview/pkg/models"
)

// Write write crontab content to file
func Write(s *models.Crontab) error {
	f, err := os.OpenFile("tmpdata", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	for _, r := range s.Records {
		if _, err = f.WriteString(fmt.Sprintf("%s %s\n", r.Schedule, r.Command)); err != nil {
			return err
		}
	}

	return nil
}
