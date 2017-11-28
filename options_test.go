package ffmpeg

import (
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
