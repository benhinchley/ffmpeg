package ffmpeg

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// stream_loop      [number]                         [input]         [X]
// itsoffset        [offset]                         [input]         [ ]
// dump_attachment  [filename]                       [input]         [ ]
// muxdelay         [seconds]                        [input]         [ ]
// muxpreload       [seconds]                        [input]         [ ]
// f                [fmt]                            [input output]  [X]
// c                [codec]                          [input output]  [ ]
// codec            [codec]                          [input output]  [ ]
// t                [duration]                       [input output]  [X]
// ss               [position]                       [input output]  [ ]
// sseof            [position]                       [input output]  [ ]
// to               [position]                       [output]        [ ]
// fs               [limit_size]                     [output]        [ ]
// timestamp        [date]                           [output]        [ ]
// metadata         [key=value]                      [output]        [ ]
// disposition      [value]                          [output]        [ ]
// target           [type]                           [output]        [ ]
// dframes          [number]                         [output]        [ ]
// frames           [framecount]                     [output]        [ ]
// q                [q]                              [output]        [ ]
// qscale           [q]                              [output]        [ ]
// filter           [filtergraph]                    [output]        [ ]
// filter_script    [filename]                       [output]        [ ]
// pre              [preset_name]                    [output]        [ ]
// attach           [filename]                       [output]        [ ]
// rc_override      [override]                       [output]        [ ]
// top              [n]                              [output]        [ ]
// shortest         []                               [output]        [ ]
// streamid         [output-stream-index:new-value]  [output]        [ ]

// WithFormat forces the the input or output file format
//
// The format is normally auto detected for input files and guessed
// from the file extension for output files, so this option is
// not needed in most cases.
func WithFormat(typ FileFormat) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-f", typ.String()}...)
		return nil
	}
}

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


func WithVideoCodec(codec Codec) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-c:v", codec.String()}...)
		return nil
	}
}

func WithAudioCodec(codec Codec) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-c:a", codec.String()}...)
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

func WithPatternType(typ string) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-pattern_type", typ}...)
		return nil
	}
}
