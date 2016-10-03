package main

import (
	"github.com/kataras/iris"
)

func ImgGetWidthCrop(ctx *iris.Context) {
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
	if !imgReal.HasVariantWithCrop(imgWidth, imgCrop) {
		// Variant does not exist
		log.Debugf("Image %s doesn't yet have a variant in file %s, trying to create it.", imgReal.Uuid, imgReal.GetVariantPathWithCrop(imgWidth, imgCrop))
		err := imgReal.MakeVariantWithCrop(imgWidth, imgCrop)
		if err != nil {
			log.Warningf("Could not create variant: %v", err)
			ctx.EmitError(iris.StatusNotFound)
			return
		}
	}
	ctx.ServeFile(imgReal.GetVariantPathWithCrop(imgWidth, imgCrop), false)
}

func ImgGetWidth(ctx *iris.Context) {
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

	if !imgReal.HasVariant(imgWidth) {
		// Variant does not exist
		log.Debugf("Image %s doesn't yet have a variant in file %s, trying to create it.", imgReal.Uuid, imgReal.GetVariantPath(imgWidth))
		err := imgReal.MakeVariant(imgWidth)
		if err != nil {
			log.Warningf("Could not create variant: %v", err)
			ctx.EmitError(iris.StatusNotFound)
			return
		}
	}
	ctx.ServeFile(imgReal.GetVariantPath(imgWidth), false)
}

func ImgGet(ctx *iris.Context) {
	// Initialize and validate UUID (image name)
	imgReal, err := ImgConf.ExistingFromUUID(ctx.Param("uuid"))
	if err != nil {
		log.Warningf("Requested image %s doesn't exist or its name is not a valid UUID v4", ctx.Param("uuid"))
		ctx.EmitError(iris.StatusNotFound)
		return
	}
	ctx.ServeFile(imgReal.GetOriginalPath(), false)
}
