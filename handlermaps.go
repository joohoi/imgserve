package main

import (
	"github.com/kataras/iris"
	"imgserve/img"
)

func GetHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){
		"/txt": img.TxtGet,
		"/jpg": img.JpgGet,
	}
}

func PostHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){
		"/txt": img.TxtPost,
	}
}
