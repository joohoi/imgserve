package img

import "github.com/kataras/iris"

func TxtGet(ctx *iris.Context) {
	ctx.ServeFile("/tmp/iris/file.txt", true)
}

func TxtPost(ctx *iris.Context) {
	ctx.Write("POST to txt")
}

func JpgGet(ctx *iris.Context) {
	ctx.ServeFile("/tmp/iris/file.jpg", false)
}
