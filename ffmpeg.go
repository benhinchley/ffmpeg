// Package ffmpeg provides a library interface to the ffmpeg command
// line interface.
package ffmpeg

import (
	"os/exec"

	multierror "github.com/hashicorp/go-multierror"
)

// Command creates a new exec.Cmd instance
func Command(global GlobalOptions, files ...*File) (*exec.Cmd, error) {
	var i, o []*File
	var err *multierror.Error

	for _, file := range files {
		switch file.typ {
		case fileTypeInput:
			if file.err != nil {
				err = multierror.Append(err, file.err)
			} else {
				i = append(i, file)
			}
		case fileTypeOutput:
			if file.err != nil {
				err = multierror.Append(err, file.err)
			} else {
				o = append(o, file)
			}
		}
	}

	if err.ErrorOrNil() != nil {
		return nil, err
	}

	var r []string
	if global != nil {
		r = append(r, global.Flags()...)
	}
	for _, input := range i {
		r = append(r, input.Flags()...)
	}
	for _, ouput := range o {
		r = append(r, ouput.Flags()...)
	}

	args := []string{"-hide_banner"}
	args = append(args, r...)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Env = append(cmd.Env, "AV_LOG_FORCE_NOCOLOR=TRUE")

	return cmd, nil
}

// GlobalOption ...
type GlobalOption func([]string) error

// FileOption ...
type FileOption func(*File) error