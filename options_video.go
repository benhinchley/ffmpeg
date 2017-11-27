package ffmpeg

import (
	"fmt"
	"strconv"
)

// Video options:
// -vframes number     set the number of video frames to output
// -r rate             set frame rate (Hz value, fraction or abbreviation)
// -s size             set frame size (WxH or abbreviation)
// -aspect aspect      set aspect ratio (4:3, 16:9 or 1.3333, 1.7777)
// -bits_per_raw_sample number  set the number of bits per raw sample
// -vn                 disable video
// -vcodec codec       force video codec ('copy' to copy stream)
// -timecode hh:mm:ss[:;.]ff  set initial TimeCode value.
// -pass n             select the pass number (1 to 3)
// -vf filter_graph    set video filters
// -ab bitrate         audio bitrate (please use -b:a)
// -b bitrate          video bitrate (please use -b:v)
// -dn                 disable data

func WithSize(w, h int) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-s", fmt.Sprintf("%dx%d", w, h)}...)
		return nil
	}
}

func WithFramerate(rate int) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-framerate", strconv.Itoa(rate)}...)
		return nil
	}
}

func WithPixelFormat(fmt PixelFormat) FileOption {
	return func(f *File) error {
		f.options = append(f.options, []string{"-pix_fmt", fmt.String()}...)
		return nil
	}
}
