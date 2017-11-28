package ffmpeg

import (
	"strings"
	"testing"
	"time"
)

func TestWithDuration(t *testing.T) {
	tests := []struct {
		Duration string
		Expected string
	}{
		{Duration: "1h20m4s", Expected: "01:20:4.000000"},
		{Duration: "1h20m4s53ms", Expected: "01:20:4.053000"},
		{Duration: "4s100ms", Expected: "00:00:4.100000"},
	}

	for _, test := range tests {
		f := &File{}
		d, _ := time.ParseDuration(test.Duration)
		opt := WithDuration(d)
		if err := opt(f); err != nil {
			t.Errorf("unable to apply option: %v", err)
		}

		if f.options[1] != test.Expected {
			t.Errorf("Expected %s got %s", test.Expected, f.options[1])
		}
	}
}

func TestWithCodec(t *testing.T) {
	tests := []struct {
		Input    []interface{}
		Expected string
	}{
		{Input: []interface{}{AllStreamSpecifier(), CodecH264}, Expected: "-c: h264"},
		{Input: []interface{}{StreamIndexSpecifier(2), CodecH264}, Expected: "-c:2 h264"},
		{Input: []interface{}{VideoStreamSpecifier(0), CodecH264}, Expected: "-c:v:0 h264"},
		{Input: []interface{}{VideoStreamSpecifier(1), CodecVp9}, Expected: "-c:v:1 vp9"},
		{Input: []interface{}{AudioStreamSpecifier(-1), CodecAac}, Expected: "-c:a aac"},
	}

	for _, test := range tests {
		f := &File{}
		opt := WithCodec(test.Input[0].(StreamSpecifier), test.Input[1].(Codec))
		if err := opt(f); err != nil {
			t.Errorf("unable to apply option: %v", err)
		}

		if strings.Join(f.options, " ") != test.Expected {
			t.Errorf("Expected %s got %s %s", test.Expected, f.options[0], f.options[1])
		}
	}
}
