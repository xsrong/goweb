package controllers

import (
	"fmt"
	"goweb/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type UsersController struct {
	Session *sessions.Session
}

func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
	middleware := func(ctx iris.Context) {
		ctx.Application().Logger()
		ctx.Next()
	}
	b.Handle("POST", "/users/new", "Create", middleware)
	b.Handle("GET", "/users/{id:int}", "Show", middleware)
	b.Handle("GET", "/get_session", "GetSession", middleware)
}

func (c *UsersController) Create(ctx iris.Context) (user models.User, err error) {
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = user.Create()
	return
}

func (c *UsersController) Show(id int) (user models.User, err error) {
	user, err = models.FindUserByID(id)
	return
}

func (c *UsersController) Login(ctx iris.Context) (user models.User, err error) {
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	if ok, err := user.Authenticate(); !ok {
		return user, err
	}
	c.Session.Set("userID", user.ID)
	return
}

func (c *UsersController) GetSession() {
	fmt.Println(c.Session.Lifetime)
}
