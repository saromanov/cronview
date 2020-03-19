package files

import (
	"fmt"
	"bufio"
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
	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 || b[0] == byte('#') {
			continue
		}
		if err := s.Err(); err != nil {
			return nil, err
		}
		fmt.Println(string(b))
	}
	return &models.Crontab{}, nil
}
