package commands

import (
	"bufio"
	"io"
	"os"
	"os/exec"
)

// ExecAndWrite provides execution of the command and writing results to the file
func ExecAndWrite(fileName string) error {
	cmd := exec.Command("crontab", "-l")

	// open the out file for writing
	outfile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(outfile)
	defer writer.Flush()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()
	return nil
}
