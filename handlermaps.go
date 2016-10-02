package main

import (
	"github.com/kataras/iris"
)

func GetHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){
		"/txt": TxtGet,
		"/img/:uuid/:width/:crop": TxtGetFull,
		"/img/:uuid/:width":       TxtGetNoCrop,
		"/img/:uuid":              TxtGetOriginal,
		"/jpg":                    JpgGet,
	}
}

func PostHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){
		"/txt": TxtPost,
	}
}
