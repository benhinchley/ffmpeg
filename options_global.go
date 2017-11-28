package ffmpeg

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// cpuflags                [flags]        [global]  [ ]
// opencl_options          [options]      [global]  [X]
// y                       []             [global]  [X]
// n                       []             [global]  [X]
// filter_threads          [nb_threads]   [global]  [X]
// stats                   []             [global]  [ ]
// progress                [url]          [global]  [ ]
// debug_ts                []             [global]  [ ]
// qphist                  []             [global]  [ ]
// benchmark               []             [global]  [ ]
// benchmark_all           []             [global]  [ ]
// timelimit               [duration]     [global]  [X]
// dump                    []             [global]  [ ]
// hex                     []             [global]  [ ]
// filter_complex          [filtergraph]  [global]  [ ]
// filter_complex_threads  [nb_threads]   [global]  [X]
// lavfi                   [filtergraph]  [global]  [ ]
// filter_complex_script   [filename]     [global]  [X]
// override_ffserver       []             [global]  [ ]
// sdp_file                [file]         [global]  [ ]
// abort_on                [flags]        [global]  [ ]

// GlobalOption configures how ffmpeg runs overall
type GlobalOption func() ([]string, error)

// GlobalOptions represents a slice of GlobalOption to be applied
type GlobalOptions []GlobalOption

// Flags generates the ffmpeg flags to be applied
func (g GlobalOptions) Flags() []string {
	// TODO: capture potential errors and return *mutlierror.Error
	var f []string
	for _, opt := range g {
		gf, _ := opt()
		f = append(f, gf...)
	}
	return f
}

// WithLogLevel sets the logging level used by ffmpeg
func WithLogLevel(l LogLevel) GlobalOption {
	return func() ([]string, error) {
		return []string{"-loglevel", l.String()}, nil
	}
}

// LogLevel ...
type LogLevel int

// Log level definitions
const (
	LogLevelQuiet   LogLevel = iota - 1 // Show nothing at all; be silent.
	LogLevelPanic   LogLevel = iota     // Only show fatal errors which could lead the process to crash, such as an assertion failure. This is not currently used for anything.
	LogLevelFatal                       // Only show fatal errors. These are errors after which the process absolutely cannot continue.
	LogLevelError                       // Show all errors, including ones which can be recovered from.
	LogLevelWarning                     // Show all warnings and errors. Any message related to possibly incorrect or unexpected events will be shown.
	LogLevelInfo                        // Show informative messages during processing. This is in addition to warnings and errors. This is the default value.
	LogLevelVerbose                     // Same as "info", except more verbose.
	LogLevelDebug                       // Show everything, including debugging information.
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

// WithOpenCLOptions sets OpenCL environment options
//
// This option is only available when FFmpeg has been compiled with "--enable-opencl"
func WithOpenCLOptions(opts map[string]string) GlobalOption {
	create := func(opts map[string]string) string {
		var f []string
		for key, value := range opts {
			f = append(f, fmt.Sprintf("%s=%s", key, value))
		}
		return strings.Join(f, ":")
	}

	return func() ([]string, error) {
		return []string{"-opencl_options", create(opts)}, nil
	}
}

// WithOverwrite sets whether to overwrite output files without asking
//
// If set to false the command will exit immediately if a specified output file already exists
func WithOverwrite(ovr bool) GlobalOption {
	return func() ([]string, error) {
		var tmp []string
		if ovr {
			tmp = append(tmp, "-y")
		} else {
			tmp = append(tmp, "-n")
		}
		return tmp, nil
	}
}

// WithNumFilterThreads defines how many threads are used to process a filter pipeline
//
// Each pipeline will produce a thread pool with this many threads
// available for parallel processing.  The default is the number of
// available CPUs.
func WithNumFilterThreads(num int) GlobalOption {
	return func() ([]string, error) {
		return []string{"-filter_threads", strconv.Itoa(num)}, nil
	}
}

// WithTimelimit sets the timelimit duration on ffmpeg.
//
// Exit after ffmpeg has been running for duration seconds.
func WithTimelimit(dur time.Duration) GlobalOption {
	create := func(dur time.Duration) string {
		matches := regexpDuration.FindAllStringSubmatch(dur.String(), -1)

		hour, _ := strconv.ParseInt(matches[0][2], 10, 32)
		minute, _ := strconv.ParseInt(matches[0][4], 10, 32)
		seconds, _ := strconv.ParseFloat(matches[0][6], 32)

		return fmt.Sprintf("%02d:%02d:%f", hour, minute, seconds)
	}
	return func() ([]string, error) {
		return []string{"-timelimit", create(dur)}, nil
	}
}

// WithNumFilterComplexThreads defines how many threads are used to process a filter_complex graph
//
// Similar to WithNumFilterThreads but used for "-filter_complex"
// graphs only. The default is the number of available CPUs.
func WithNumFilterComplexThreads(num int) GlobalOption {
	return func() ([]string, error) {
		return []string{"-filter_complex_threads", strconv.Itoa(num)}, nil
	}
}

// WithFilterComplexScript ...
func WithFilterComplexScript(filename string) GlobalOption {
	return func() ([]string, error) {
		return []string{"-filter_complex_script", filename}, nil
	}
}
