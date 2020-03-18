package files

import (
	"fmt"
	"os"
	"github.com/saromanov/cronview/pkg/models"
)

// Write write crontab content to file
func Write(m *models.Crontab) error {
	if m == nil {
		return fmt.Errorf("crontab struct is not defined")
	}
	f, err := os.OpenFile("tmpdata", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}
