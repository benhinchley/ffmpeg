package ffmpeg

import multierror "github.com/hashicorp/go-multierror"

// Input ...
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

// Output ...
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

// File ...
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
	case fileTypeOutput:
		return append(f.options, f.path)
	default:
		return f.options
	}
}

type fileType int

const (
	fileTypeGlobal fileType = iota - 1
	fileTypeInput  fileType = iota
	fileTypeOutput
)
