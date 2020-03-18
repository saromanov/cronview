package files

import (
	"os"
)

// Write write crontab content to file
func Write(records []string) error {
	f, err := os.OpenFile("tmpdata", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	for _, r := range records {
		if _, err = f.WriteString(r); err != nil {
			return err
		}
	}

	return nil
}
