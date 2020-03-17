package files

import (
	"bufio"
	"os"
	"github.com/saromanov/cronview/pkg/cmd"
	"github.com/saromanov/cronview/pkg/models"
)

// Read open a file an read content into crontab format
func Read(file string) (*hostFile, error) {
	fromFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return Read(fromFile, false)
}

// Read returns crontab file
func Read(r io.Reader, strict bool) (*models.Crontab, error) {
	c := &models.Crontab
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		ln++
		b := s.Bytes()

		switch {

		if err := s.Err(); err != nil {
			return nil, err
		}
	}
	return с, nil
}