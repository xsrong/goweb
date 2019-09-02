package controllers

import (
	"errors"
	"goweb/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type PostsController struct {
	Session *sessions.Session
}

func (c *PostsController) BeforeActivation(b mvc.BeforeActivation) {
	middleware := func(ctx iris.Context) {
		ctx.Application().Logger()
		ctx.Next()
	}

	b.Handle("POST", "/posts/new", "Create", middleware)
}

func (c *PostsController) Create(ctx iris.Context) (post models.Post, err error) {
	currentUserID, err := c.Session.GetInt("userID")
	if err != nil {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	currentUser, err := models.FindUserByID(currentUserID)
	if err != nil {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	if err = ctx.ReadJSON(&post); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = currentUser.AddPost(&post)
	return
}
