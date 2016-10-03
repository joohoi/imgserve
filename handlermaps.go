package main

import (
	"github.com/kataras/iris"
)

func GetHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){
		"/img/:uuid/:width/:crop": ImgGetWidthCrop,
		"/img/:uuid/:width":       ImgGetWidth,
		"/img/:uuid":              ImgGet,
	}
}

func PostHandlerMap() map[string]func(*iris.Context) {
	return map[string]func(*iris.Context){}
}
