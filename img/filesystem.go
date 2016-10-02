package img

import (
	"errors"
	"os"
	"strings"
)

var orig string = "original"
var delim string = "_"

func GetFullpath(path string, uuid string) (string, error) {
	fullpath := strings.Join([]string{path, "/", uuid}, "")
	if Exists(fullpath) {
		return fullpath, nil
	} else {
		return "", errors.New("File does not exist")
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetVariantPath(path string, width string, crop string) string {
	return strings.Join([]string{path, "/", orig, delim, width, delim, crop}, "")
}

func HasVariant(path string, width string, crop string) bool {
	return Exists(GetVariantPath(path, width, crop))
}
