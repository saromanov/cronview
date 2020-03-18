package commands

import (
	"fmt"
	"os/exec"
)

// ExecAndWrite provides execution of the command and writing results to the file
func ExecAndWrite() error {
	_, err := exec.Command("crontab", "-l", ">data.tmp").Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %v", err)
	}
	return nil
}
