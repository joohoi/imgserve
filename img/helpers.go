package img

import (
	"math"
	"string"
)

func CropHeight(width int, crop string) int {
	cropRatio := string.Split(crop, "x")
	return int(math.Floor(float64(width) / (float64(cropRatio[0]) / float64(cropRatio[1]))))
}
