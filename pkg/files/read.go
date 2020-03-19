package files

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/saromanov/cronview/pkg/models"
)

// Read open a file an read content into crontab format
func Read(file string) (*models.Crontab, error) {
	fromFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return read(fromFile, false)
}

// Read returns crontab file
func read(r io.Reader, strict bool) (*models.Crontab, error) {
	s := bufio.NewScanner(r)
	cronTab := &models.Crontab{}
	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 || b[0] == byte('#') {
			continue
		}
		if err := s.Err(); err != nil {
			return nil, err
		}
		idx, schedule := getCronSchedule(b)
		if idx == 0 || schedule == "" {
			return nil, fmt.Errorf("unable to get cron schedule")
		}
		cronTab.Records = append(cronTab.Records, models.Record{
			Schedule: schedule,
			Command:  string(b[idx : len(b)-1]),
		})
	}
	return cronTab, nil
}

// return cron schedule based of cron task line
// its return index from beginning of the command
// and raw cron schedule
func getCronSchedule(data []byte) (int, string) {
	spaces := 0
	for i, d := range data {
		if d == byte(' ') {
			spaces++
		}
		if spaces == 5 {
			return i, string(data[0:i])
		}
	}
	return 0, ""
}
