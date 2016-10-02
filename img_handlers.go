package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func TxtGet(ctx *iris.Context) {
	ctx.ServeFile("/tmp/iris/file.txt", true)
}

func TxtPost(ctx *iris.Context) {
	ctx.Write("POST to txt")
}

func JpgGet(ctx *iris.Context) {
	ctx.ServeFile("/tmp/iris/file.jpg", false)
}

func TxtGetFull(ctx *iris.Context) {
	// Initialize and validate UUID (image name)
	imgReal, err := ImgConf.ExistingFromUUID(ctx.Param("uuid"))
	if err != nil {
		log.Warningf("Requested image %s doesn't exist or its name is not a valid UUID v4", ctx.Param("uuid"))
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	// Initialize and validate requested width
	imgWidth := ctx.Param("width")
	if !ImgConf.ValidWidth(imgWidth) {
		log.Warningf("Requested width %s not found for image %s", imgWidth, imgReal.Uuid)
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	// Initialize and validate requested crop
	imgCrop := ctx.Param("crop")
	if !ImgConf.ValidCrop(imgCrop) {
		log.Warningf("Requested crop %s not found for image %s", imgCrop, imgReal.Uuid)
		ctx.EmitError(iris.StatusNotFound)
		return
	}
	if !imgReal.HasVariant(imgWidth, imgCrop) {
		// Variant does not exist
		log.Debugf("Image %s doesn't yet have a variant in file %s", imgReal.Uuid, imgReal.GetVariantPath(imgWidth, imgCrop))
		ctx.EmitError(iris.StatusNotFound)
		return
	}
	ctx.ServeFile(imgReal.GetVariantPath(imgWidth, imgCrop), false)
}

func TxtGetNoCrop(ctx *iris.Context) {
	imgUUID := ctx.Param("uuid")
	imgWidth := ctx.Param("width")
	retString := fmt.Sprintf("NoCrop: This is image %s with width %s and crop %s", imgUUID, imgWidth)
	ctx.Write(retString)
}

func TxtGetOriginal(ctx *iris.Context) {
	imgUUID := ctx.Param("uuid")
	retString := fmt.Sprintf("Original: This is image %s", imgUUID)
	ctx.Write(retString)
}
