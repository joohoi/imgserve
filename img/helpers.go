package img

import (
	"math"
	"strconv"
	"strings"
)

func CropHeight(width int, crop string) (int, error) {
	cropRatio := strings.Split(crop, "x")
	cropX, err := strconv.ParseInt(cropRatio[0], 10, 0)
	if err != nil {
		return 0, nil
	}
	cropY, err := strconv.ParseInt(cropRatio[1], 10, 0)
	if err != nil {
		return 0, nil
	}
	return int(math.Floor(float64(width) / (float64(cropX) / float64(cropY)))), nil
}
