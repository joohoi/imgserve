package img

import (
	"github.com/disintegration/imaging"
	"image"
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
	var newVariant *image.NRGBA
	data, err := imaging.Open(i.GetOriginalPath())
	if err != nil {
		return err
	}
	crop_widthInt, err := strconv.ParseInt(width, 10, 0)
	if err != nil {
		return err
	}
	crop_width := int(crop_widthInt)
	bounds := data.Bounds()
	orig_width := bounds.Max.X
	orig_height := bounds.Max.Y
	crop_height, err := CropHeight(crop_width, crop)
	if err != nil {
		return err
	}
	// Get the actual crop values
	if (orig_width < crop_width) || (orig_height < crop_height) {
		newVariant = imaging.Fill(data, crop_width, crop_height, imaging.Center, imaging.Lanczos)
	} else {
		// TODO: get user specified crop size from database
		// and use Crop(img image.Image, rect image.Rectangle) *image.NRGBA
		newVariant = imaging.CropAnchor(data, crop_width, crop_height, imaging.Center)
	}
	err = imaging.Save(newVariant, i.GetVariantPathWithCrop(width, crop))
	if err != nil {
		return err
	}
	return nil
}
