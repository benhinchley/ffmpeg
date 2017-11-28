package ffmpeg

import multierror "github.com/hashicorp/go-multierror"

// Input creates a new File instance that represents an input file
func Input(path string, opts ...FileOption) *File {
	f := &File{
		path: path,
		typ:  fileTypeInput,
	}

	var errs *multierror.Error
	for _, opt := range opts {
		if err := opt(f); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	f.err = errs.ErrorOrNil()

	return f
}

// Output creates a new File instance that represents an output file
func Output(path string, opts ...FileOption) *File {
	f := &File{
		path: path,
		typ:  fileTypeOutput,
	}

	var errs *multierror.Error
	for _, opt := range opts {
		if err := opt(f); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	f.err = errs.ErrorOrNil()

	return f
}

// File represents an ffmpeg input or output file
type File struct {
	path    string
	options []string
	typ     fileType
	err     error
}

// Flags generates the ffmpeg flags for the specified file
func (f *File) Flags() []string {
	switch f.typ {
	case fileTypeInput:
		return append(f.options, []string{"-i", f.path}...)
	default:
		return append(f.options, f.path)
	}
}

type fileType int

const (
	fileTypeInput fileType = iota
	fileTypeOutput
)

// FileOption configures how a file is handled by ffmpeg
type FileOption func(*File) error
