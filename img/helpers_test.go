package img

import (
	"testing"
)

func TestCropHeight(t *testing.T) {
	for i, test := range []struct {
		in_width  int
		in_crop   string
		output    int
		shouldErr bool
	}{
		{800, "2x1", 400, false},
		{1600, "16x9", 900, false},
		{600, "notratio", 0, true},
		{100, "x2", 0, true},
		{100, "2x", 0, true},
		{100, "1xx", 0, true},
		{100, "xx2", 0, true},
	} {
		ret, err := CropHeight(test.in_width, test.in_crop)
		if test.shouldErr && err != nil {
			t.Errorf("Test %d: Expected error, but there wasn't any", i)
		}
		if !test.shouldErr && err != nil {
			t.Errorf("Test %d: Expected no error, but there was one: %v", i, err)
		}
		if ret != test.output {
			t.Errorf("Test %d: Expected value %d, but got %d", i, test.output, ret)
		}
	}
}
