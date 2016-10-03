package img

import (
	"github.com/disintegration/imaging"
	"strconv"
)

func (i Img) MakeVariant(width string) error {
	data, err := imaging.Open(i.GetOriginalPath())
	if err != nil {
		return err
	}
	widthInt, err := strconv.ParseInt(width, 10, 0)
	if err != nil {
		return err
	}
	newVariant := imaging.Resize(data, int(widthInt), 0, imaging.Lanczos)
	err = imaging.Save(newVariant, i.GetVariantPath(width))
	if err != nil {
		return err
	}
	return nil
}

func (i Img) MakeVariantWithCrop(width string, crop string) error {
	data, err := imaging.Open(i.GetOriginalPath())
	if err != nil {
		return err
	}
	// Get the actual crop values
	newVariant := imaging.Fill(data, 100, 100, imaging.Center, imaging.Lanczos)
	err = imaging.Save(newVariant, i.GetVariantPathWithCrop(width, crop))
	if err != nil {
		return err
	}
	return nil
}
