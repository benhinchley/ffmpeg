package ffmpeg

import "fmt"

type StreamType int

const (
	StreamTypeAll   StreamType = iota - 1
	StreamTypeVideo StreamType = iota
	StreamTypeAudio
	StreamTypeSubtitle
	StreamTypeData
	StreamTypeAttachment
)

func (st StreamType) String() string {
	switch st {
	case StreamTypeVideo:
		return "v"
	case StreamTypeAudio:
		return "a"
	case StreamTypeSubtitle:
		return "s"
	case StreamTypeData:
		return "d"
	case StreamTypeAttachment:
		return "t"
	default:
		return ""
	}
}

type StreamSpecifier struct {
	Stream StreamType
	Idx    int
}

func (s StreamSpecifier) String() string {
	if s.Idx == -1 {
		return ":" + s.Stream.String()
	} else if s.Stream == StreamTypeAll {
		return fmt.Sprintf(":%d", s.Idx)
	}

	return fmt.Sprintf(":%s:%d", s.Stream, s.Idx)
}

func AllStreamSpecifier() StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeAll,
		Idx:    -1,
	}
}

func StreamIndexSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeAll,
		Idx:    idx,
	}
}

func VideoStreamSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeVideo,
		Idx:    idx,
	}
}

func AudioStreamSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeAudio,
		Idx:    idx,
	}
}

func SubtitleStreamSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeSubtitle,
		Idx:    idx,
	}
}

func DataStreamSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeData,
		Idx:    idx,
	}
}

func AttachmentStreamSpecifier(idx int) StreamSpecifier {
	return StreamSpecifier{
		Stream: StreamTypeAttachment,
		Idx:    idx,
	}
}
