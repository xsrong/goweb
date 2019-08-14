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
