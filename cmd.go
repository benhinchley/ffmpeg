package ffmpeg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Cmd represents an ffmpeg command being prepared or run.
type Cmd struct {
	Args []string

	cmd *exec.Cmd
}

// Run starts the specified command and waits for it to complete.
func (cmd *Cmd) Run() error {
	var stderr bytes.Buffer
	cmd.cmd.Stderr = &stderr
	err := cmd.cmd.Run()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			fmt.Fprintf(os.Stderr, "%s\n", stderr.String())
			return err
		default:
			return err
		}
	}
	return nil
}
