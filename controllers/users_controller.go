package controllers

import (
	"errors"
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
	b.Handle("POST", "/users/{id:int}/edit", "Update", middleware)
	b.Handle("POST", "/login", "Login", middleware)
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
	// ctx.Header("Access-Control-Allow-Origin", "*")
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

func (c *UsersController) Update(id int, ctx iris.Context) (user models.User, err error) {
	userID, err := c.Session.GetInt("userID")
	if err != nil || userID != id {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	params := map[string]string{}
	if err = ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	user, err = models.FindUserByID(id)
	if err != nil {
		return
	}
	err = user.Update(params)
	return
}
