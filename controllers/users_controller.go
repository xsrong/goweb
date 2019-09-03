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
	b.Handle("PATCH", "/users/{id:int}/edit", "Update", middleware)
	b.Handle("POST", "/login", "Login", middleware)
	b.Handle("DELETE", "/logout", "Logout", middleware)
	b.Handle("POST", "/users/{id:int}/follow/{followedID:int}", "Follow", middleware)
	b.Handle("DELETE", "/users/{id:int}/unfollow/{unfollowID:int}", "Unfollow", middleware)
	b.Handle("GET", "/users/{id:int}/following", "Following", middleware)
	b.Handle("GET", "/users/{id:int}/followers", "Followers", middleware)
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
	if IsLoggedIn(c.Session) {
		err = errors.New("logged in already")
		return
	}
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

func (c *UsersController) Logout() (err error) {
	if !IsLoggedIn(c.Session) {
		err = errors.New("not logged in")
		return
	}
	if !c.Session.Delete("userID") {
		err = errors.New("logout failed. please try again later")
	}
	return
}

func (c *UsersController) Update(id int, ctx iris.Context) (user models.User, err error) {
	if !IsLoggedIn(c.Session) || !IsCurrentUser(id, c.Session) {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	params := map[string]string{}
	if err = ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	user = models.User{ID: id}
	err = user.Update(params)
	return
}

func (c *UsersController) Follow(id, followedID int) (err error) {
	if !IsLoggedIn(c.Session) || !IsCurrentUser(id, c.Session) {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	if id == followedID {
		err = errors.New("cannot follow yourself")
		return
	}
	fromUser := models.User{ID: id}
	toUser := models.User{ID: followedID}
	err = fromUser.Follow(toUser)
	return
}

func (c *UsersController) Unfollow(id, unfollowID int) (err error) {
	if !IsLoggedIn(c.Session) || !IsCurrentUser(id, c.Session) {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	if id == unfollowID {
		err = errors.New("cannot unfollow yourself")
		return
	}
	fromUser := models.User{ID: id}
	toUser := models.User{ID: unfollowID}
	err = fromUser.Unfollow(toUser)
	return
}

func (c *UsersController) Following(id int) (users []models.User, err error) {
	if !IsLoggedIn(c.Session) || !IsCurrentUser(id, c.Session) {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	user := models.User{ID: id}
	users, err = user.Followings()
	return
}

func (c *UsersController) Followers(id int) (users []models.User, err error) {
	if !IsLoggedIn(c.Session) || !IsCurrentUser(id, c.Session) {
		err = errors.New("authenticate failed! please login and try again")
		return
	}
	user := models.User{ID: id}
	users, err = user.Followers()
	return
}
