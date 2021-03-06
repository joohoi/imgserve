package img

import (
	"errors"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

type ImgConfig struct {
	Path   string
	Widths []int64
	Crops  []string
}

func (i ImgConfig) NewImg() Img {
	return Img{
		Uuid: uuid.NewV4(),
		path: i.Path,
	}
}

// TODO: remove this and make assumption the file exists
func (i ImgConfig) ExistingFromUUID(imgUuid string) (Img, error) {
	realUuid, err := uuid.FromString(imgUuid)
	if err != nil {
		return Img{}, err
	}
	retImg := Img{
		Uuid: realUuid,
		path: strings.Join([]string{i.Path, "/", realUuid.String()}, ""),
	}
	if !retImg.Exists() {
		return Img{}, errors.New("File does not exist")
	}
	return retImg, nil
}

func (i ImgConfig) ValidWidth(width string) bool {
	widthInt, err := strconv.ParseInt(width, 10, 0)
	if err != nil {
		return false
	}
	if len(i.Widths) > 0 {
		for _, v := range i.Widths {
			if v == widthInt {
				return true
			}
		}
	}
	return false
}

func (i ImgConfig) ValidCrop(crop string) bool {
	if len(i.Crops) > 0 {
		for _, v := range i.Crops {
			if v == crop {
				return true
			}
		}
	}
	return false
}
