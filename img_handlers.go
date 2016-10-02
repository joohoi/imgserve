package main

import (
	"fmt"
	"github.com/kataras/iris"
	"imgserve/img"
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
	imgUUID := ctx.Param("uuid")
	imgWidth := ctx.Param("width")
	imgCrop := ctx.Param("crop")
	imgFullPath, err := img.GetFullpath(Conf.Path, imgUUID)
	if err != nil {
		// Does not exist
		ctx.EmitError(iris.StatusNotFound)
		return
	}
	if !img.HasVariant(imgFullPath, imgWidth, imgCrop) {
		// Variant does not exist
		ctx.EmitError(iris.StatusNotFound)
		return
	}
	ctx.ServeFile(img.GetVariantPath(imgFullPath, imgWidth, imgCrop), false)
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
