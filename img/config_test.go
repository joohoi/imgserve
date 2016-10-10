package img

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadConfig(t *testing.T) {
	for i, test := range []struct {
		in_file   []byte
		output    ImgConfig
		shouldErr bool
	}{
		{
			[]byte("path = \"/tmp/imgserve\"\nwidths = [480, 600, 768]\ncrops = [\"1x1\", \"16x9\"]"),
			ImgConfig{
				Path:   "/tmp/imgserve",
				Widths: []int64{480, 600, 768},
				Crops:  []string{"1x1", "16x9"},
			},
			false,
		},
	} {
		tmpfile, err := ioutil.TempFile("", "imgtest")
		if err != nil {
			t.Error("Could not create temporary file")
		}
		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write(test.in_file); err != nil {
			t.Error("Could not write to temporary file")
		}

		if err := tmpfile.Close(); err != nil {
			t.Error("Could not close temporary file")
		}
		// write
		ret, err := ReadConfig(tmpfile.Name())
		if test.shouldErr && err != nil {
			t.Errorf("Test %d: Expected error, but there wasn't any", i)
		}
		if !test.shouldErr && err != nil {
			t.Errorf("Test %d: Expected no error, but there was one: %v", i, err)
		}
		if ret.Path != test.output.Path {
			t.Errorf("Test %d: Expected path %s, but got %s", i, test.output.Path, ret.Path)
		}
		if !reflect.DeepEqual(ret.Widths, test.output.Widths) {
			t.Errorf("Test %d: Expected widths: %v, but got %v", i, test.output.Widths, ret.Widths)
		}
		if !reflect.DeepEqual(ret.Crops, test.output.Crops) {
			t.Errorf("Test %d: Expected crops: %v, but got %v", i, test.output.Crops, ret.Crops)
		}
	}
}
