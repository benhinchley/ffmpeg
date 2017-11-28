package ffmpeg

import "testing"

func TestStreamSpecifier(t *testing.T) {
	tests := []struct {
		Specifier StreamSpecifier
		Expected  string
	}{
		{Specifier: StreamSpecifier{
			Stream: StreamTypeAll,
			Idx:    -1,
		}, Expected: ":"},
		{Specifier: StreamSpecifier{
			Stream: StreamTypeAll,
		}, Expected: ":0"},
		{Specifier: StreamSpecifier{
			Stream: StreamTypeVideo,
			Idx:    -1,
		}, Expected: ":v"},
		{Specifier: StreamSpecifier{
			Stream: StreamTypeVideo,
			Idx:    10,
		}, Expected: ":v:10"},
		{Specifier: AllStreamSpecifier(), Expected: ":"},
		{Specifier: VideoStreamSpecifier(-1), Expected: ":v"},
		{Specifier: AudioStreamSpecifier(-1), Expected: ":a"},
		{Specifier: DataStreamSpecifier(-1), Expected: ":d"},
		{Specifier: AttachmentStreamSpecifier(-1), Expected: ":t"},
		{Specifier: StreamIndexSpecifier(10), Expected: ":10"},
		{Specifier: VideoStreamSpecifier(11), Expected: ":v:11"},
		{Specifier: AudioStreamSpecifier(0), Expected: ":a:0"},
		{Specifier: DataStreamSpecifier(1), Expected: ":d:1"},
		{Specifier: AttachmentStreamSpecifier(5), Expected: ":t:5"},
	}

	for _, test := range tests {
		if test.Specifier.String() != test.Expected {
			t.Errorf("Expected %s got %s", test.Expected, test.Specifier.String())
		}
	}
}
