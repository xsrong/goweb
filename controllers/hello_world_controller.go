package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	// "github.com/kataras/iris/sessions"
)

type HelloWorldController struct {
	// Session *sessions.Session
}

func (c *HelloWorldController) BeforeActivation(b mvc.BeforeActivation) {
	middleware := func(ctx iris.Context) {
		ctx.Application().Logger()
		ctx.Next()
	}
	b.Handle("GET", "/home", "Home", middleware)
}

func (c *HelloWorldController) Home() (mes string, err error, statusCode int) {
	return "Hello World!", nil, 200
}
