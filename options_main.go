package ffmpeg

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// FLAG             SPEC   ARGS                              AFFECTS        IMPL
// stream_loop      false  [number]                         [input]         [X]
// itsoffset        false  [offset]                         [input]         [ ]
// dump_attachment  true   [filename]                       [input]         [ ]
// muxdelay         false  [seconds]                        [input]         [ ]
// muxpreload       false  [seconds]                        [input]         [ ]
// f                false  [fmt]                            [input output]  [X]
// c                true   [codec]                          [input output]  [X]
// codec            true   [codec]                          [input output]  [X]
// t                false  [duration]                       [input output]  [X]
// ss               false  [position]                       [input output]  [ ]
// sseof            false  [position]                       [input output]  [ ]
// to               false  [position]                       [output]        [ ]
// fs               false  [limit_size]                     [output]        [ ]
// timestamp        false  [date]                           [output]        [ ]
// metadata         true   [key=value]                      [output]        [ ]
// disposition      true   [value]                          [output]        [ ]
// target           false  [type]                           [output]        [ ]
// dframes          false  [number]                         [output]        [ ]
// frames           true   [framecount]                     [output]        [ ]
// q                true   [q]                              [output]        [ ]
// qscale           true   [q]                              [output]        [ ]
// filter           true   [filtergraph]                    [output]        [ ]
// filter_script    true   [filename]                       [output]        [ ]
// pre              true   [preset_name]                    [output]        [ ]
// attach           false  [filename]                       [output]        [ ]
// rc_override      true   [override]                       [output]        [ ]
// top              true   [n]                              [output]        [ ]
// shortest         false  []                               [output]        [ ]
// streamid         false  [output-stream-index:new-value]  [output]        [ ]

// WithStreamLoop sets the number of times input stream shall be looped
//
// loop 0 means no loop, loop -1 means infinite loop
func WithStreamLoop(loop int) FileOption {
	return func(f *File) error {
		if f.typ == fileTypeInput {
			f.options = append(f.options, []string{"-stream_loop", strconv.Itoa(loop)}...)
		} else {
			return fmt.Errorf("unable to apply -stream_loop flag: not input file")
		}
		return nil
	}
}

// WithFormat forces the the input or output file format
//
// The format is normally auto detected for input files and guessed
// from the file extension for output files, so this option is
// not needed in most cases.
func WithFormat(ff FileFormat) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-f", ff.String()}...)
		return nil
	}
}

// WithCodec selects an encoder (when used before an output file)
// or a decoder (when used before an input file) for one or more streams
func WithCodec(stream StreamSpecifier, codec Codec) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-c" + stream.String(), codec.String()}...)
		return nil
	}
}

var regexpDuration = regexp.MustCompile(`((\d+)h)?((\d+)m)?(([0-9.]+)s)?`)

// WithDuration when used as an input option limits the duration of the data read from the input file
// and when used as an output option limits stops writing the ouput after its duration reaches duration.
func WithDuration(dur time.Duration) FileOption {
	create := func(dur time.Duration) string {
		matches := regexpDuration.FindAllStringSubmatch(dur.String(), -1)

		hour, _ := strconv.ParseInt(matches[0][2], 10, 32)
		minute, _ := strconv.ParseInt(matches[0][4], 10, 32)
		seconds, _ := strconv.ParseFloat(matches[0][6], 32)

		return fmt.Sprintf("%02d:%02d:%f", hour, minute, seconds)
	}
	return func(f *File) error {
		f.options = append(f.options, []string{"-t", create(dur)}...)
		return nil
	}
}

// WithFileSizeLimit sets the file size limit, expressed in bytes
func WithFileSizeLimit(limit int) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-fs", strconv.Itoa(limit)}...)
		return nil
	}
}

// WithTimestamp sets the recording timestamp in the container
func WithTimestamp(date time.Time) FileOption {
	create := func(date time.Time) string {
		return ""
	}
	return func(f *File) error {
		if f.typ == fileTypeOutput {
			f.options = append(f.options, []string{"-timestamp", create(date)}...)
		} else {
			return fmt.Errorf("unable to apply -timestamp flag: not output file")
		}
		return nil
	}
}
