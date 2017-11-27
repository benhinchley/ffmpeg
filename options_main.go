package ffmpeg

import (
	"fmt"
	"strconv"
	"time"
)

// Per-file main options:
// -f fmt              force format
// -c codec            codec name
// -codec codec        codec name
// -pre preset         preset name
// -map_metadata outfile[,metadata]:infile[,metadata]  set metadata information of outfile from infile
// -t duration         record or transcode "duration" seconds of audio/video
// -to time_stop       record or transcode stop time
// -fs limit_size      set the limit file size in bytes
// -ss time_off        set the start time offset
// -sseof time_off     set the start time offset relative to EOF
// -seek_timestamp     enable/disable seeking by timestamp with -ss
// -timestamp time     set the recording timestamp ('now' to set the current time)
// -metadata string=string  add metadata
// -program title=string:st=number...  add program with specified streams
// -target type        specify target file type ("vcd", "svcd", "dvd", "dv" or "dv50" with optional prefixes "pal-", "ntsc-" or "film-")
// -apad               audio pad
// -frames number      set the number of frames to output
// -filter filter_graph  set stream filtergraph
// -filter_script filename  read stream filtergraph description from a file
// -reinit_filter      reinit filtergraph on input parameter changes
// -discard            discard
// -disposition        disposition

func WithFormat(fmt FileFormat) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-f", fmt.String()}...)
		return nil
	}
}

func WithOverwrite(ovr bool) FileOption {
	return func(f *File) error {
		if ovr {
			f.options = append(f.options, "-y")
		} else {
			f.options = append(f.options, "-n")
		}
		return nil
	}
}

func WithStreamLoop(loop int) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-stream_loop", strconv.Itoa(loop)}...)
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

func WithDuration(dur time.Duration) FileOption {
	create := func(dur time.Duration) string {
		return fmt.Sprintf("HH:MM:SS")
	}
	return func(f *File) error {
		f.options = append(f.options, []string{"-t", create(dur)}...)
		return nil
	}
}

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
