package controllers

import (
	"goweb/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

func CurrentUser(ctx iris.Context, sess *sessions.Session) (user models.User) {
	entry := ctx.Values().Get("user")
	if entry != nil {
		user = entry.(models.User)
	} else {
		id, _ := sess.GetInt("userID")
		if id <= 0 {
			return
		}
		user, _ = models.FindUserByID(id)
		if user.ID != 0 {
			ctx.Values().Set("user", user)
		}
	}
	return
}

func IsLoggedIn(sess *sessions.Session) bool {
	userID, _ := sess.GetInt("userID")
	return userID > 0
}

func IsCurrentUser(userID int, sess *sessions.Session) bool {
	otherID, _ := sess.GetInt("userID")
	return userID == otherID
}
