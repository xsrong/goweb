package controllers

import (
	"goweb/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type UsersController struct{}

func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
	middleware := func(ctx iris.Context) {
		ctx.Application().Logger()
		ctx.Next()
	}
	b.Handle("POST", "/users/new", "Create", middleware)
}

func (c *UsersController) Create(ctx iris.Context) (user models.User, err error) {
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = user.Create()
	return
}
