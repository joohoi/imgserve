package img

import (
	"errors"
	"github.com/satori/go.uuid"
	"image"
	"os"
	"strings"
)

type Img struct {
	Uuid     uuid.UUID                  `db:"uuid"`
	Filename string                     `db:"filename"`
	Mimetype string                     `db:"mimetype"`
	Width    int                        `db:"width"`
	Height   int                        `db:"height"`
	path     string                     `db:"path"`
	Crops    map[string]image.Rectangle `db:"crops"`
}

var orig string = "original"
var delim string = "_"

func (i Img) GetFullpath(path string, uuid string) (string, error) {
	fullpath := strings.Join([]string{path, "/", uuid}, "")
	if i.Exists() {
		return fullpath, nil
	} else {
		return "", errors.New("File does not exist")
	}
}

func (i Img) Exists() bool {
	if _, err := os.Stat(i.path); err == nil {
		return true
	}
	return false
}

func (i Img) GetOriginalPath() string {
	return strings.Join([]string{i.path, "/", orig, ".jpg"}, "")
}

func (i Img) GetVariantPath(width string) string {
	return strings.Join([]string{i.path, "/", orig, delim, width, ".jpg"}, "")
}

func (i Img) GetVariantPathWithCrop(width string, crop string) string {
	return strings.Join([]string{i.path, "/", orig, delim, width, delim, crop, ".jpg"}, "")
}

func (i Img) HasVariant(width string) bool {
	if _, err := os.Stat(i.GetVariantPath(width)); err == nil {
		return true
	}
	return false
}

func (i Img) HasVariantWithCrop(width string, crop string) bool {
	if _, err := os.Stat(i.GetVariantPathWithCrop(width, crop)); err == nil {
		return true
	}
	return false
}
