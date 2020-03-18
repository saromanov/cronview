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
	for s.Scan() {
		b := s.Bytes()
		fmt.Println(b)
		if err := s.Err(); err != nil {
			return nil, err
		}
	}
	return &models.Crontab{}, nil
}
