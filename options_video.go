package ffmpeg

import (
	"fmt"
)

// FLAG              SPEC   ARGS              AFFECTS         IMPL
// hwaccel           true   [hwaccel]         [input]         [ ]
// hwaccel_device    true   [hwaccel_device]  [input]         [ ]
// r                 true   [fps]             [input output]  [ ]
// s                 true   [size]            [input output]  [X]
// pix_fmt           true   [format]          [input output]  [X]
// sws_flags         false  [flags]           [input output]  [ ]
// vframes           false  [number]          [output]        [ ]
// aspect            true   [aspect]          [output]        [ ]
// vn                false  []                [output]        [ ]
// vcodec            false  [codec]           [output]        [ ]
// pass              true   [n]               [output]        [ ]
// passlogfile       true   [prefix]          [output]        [ ]
// vf                false  [filtergraph]     [output]        [ ]
// vtag              false  [fourcc/tag]      [output]        [ ]
// force_key_frames  true   [time[,time...]]  [output]        [ ]
// force_key_frames  true   [expr:expr]       [output]        [ ]
// copyinkf          true   []                [output]        [ ]

func WithSize(stream StreamSpecifier, w, h int) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-s" + stream.String(), fmt.Sprintf("%dx%d", w, h)}...)
		return nil
	}
}

func WithPixelFormat(stream StreamSpecifier, pf PixelFormat) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-pix_fmt" + stream.String(), pf.String()}...)
		return nil
	}
}
