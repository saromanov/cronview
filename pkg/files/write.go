package files

import "os"

// WriteToFile write crontab content to file
func WriteToFile(f *os.File) error {
	err := f.Truncate(0)
	if err != nil {
		return err
	}

	return nil
}