package ffmpeg

import (
	"fmt"
	"strconv"
	"strings"
)

// Global options (affect whole program instead of just one file:
// -loglevel loglevel  set logging level
// -v loglevel         set logging level
// -report             generate a report
// -max_alloc bytes    set maximum size of a single allocated block
// -opencl_options     set OpenCL environment options
// -ignore_unknown     Ignore unknown stream types
// -filter_threads     number of non-complex filter threads
// -filter_complex_threads  number of threads for -filter_complex
// -stats              print progress report during encoding
// -max_error_rate ratio of errors (0.0: no errors, 1.0: 100% error  maximum error rate
// -bits_per_raw_sample number  set the number of bits per raw sample
// -vol volume         change audio volume (256=normal)

type GlobalOptions []GlobalOption

func (g GlobalOptions) Flags() []string {
	var f []string
	for _, opt := range g {
		opt(f)
	}
	return f
}

func WithLogLevel(l LogLevel) GlobalOption {
	return func(f []string) error {
		f = append(f, []string{"-loglevel", l.String()}...)
		return nil
	}
}

type LogLevel int

const (
	LogLevelQuiet LogLevel = iota - 1
	LogLevelPanic LogLevel = iota
	LogLevelFatal
	LogLevelError
	LogLevelWarning
	LogLevelInfo
	LogLevelVerbose
	LogLevelDebug
	LogLevelTrace
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelPanic:
		return "panic"
	case LogLevelFatal:
		return "fatal"
	case LogLevelError:
		return "error"
	case LogLevelWarning:
		return "warning"
	case LogLevelInfo:
		return "info"
	case LogLevelVerbose:
		return "verbose"
	case LogLevelDebug:
		return "debug"
	case LogLevelTrace:
		return "trace"
	default:
		return "quiet"
	}
}

func WithMaxAlloc(max int) GlobalOption {
	return func(f []string) error {
		f = append(f, []string{"-max_alloc", strconv.Itoa(max)}...)
		return nil
	}
}

func WithOpenCLOptions(opts map[string]string) GlobalOption {
	create := func(opts map[string]string) string {
		var f []string
		for key, value := range opts {
			f = append(f, fmt.Sprintf("%s=%s", key, value))
		}
		return strings.Join(f, ":")
	}

	return func(f []string) error {
		f = append(f, []string{"-opencl_options", create(opts)}...)
		return nil
	}
}
