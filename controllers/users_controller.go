package controllers

import (
	"goweb/models"
	"strconv"

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
	b.Handle("POST", "/login", "Login", middleware)
	b.Handle("GET", "/get_id", "GetCurrentUserID", middleware)
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
	if user, err = user.Authenticate(); err != nil {
		return
	}
	c.Session.Set("userID", user.ID)
	return
}

func (c *UsersController) GetCurrentUserID(ctx iris.Context) string {
	user := CurrentUser(ctx, c.Session)
	return "the user id is " + strconv.Itoa(user.ID)
}
