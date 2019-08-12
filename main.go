package main

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"

	"goweb/controllers"
)

func main() {
	app := weiboApp()
	app.Run(iris.Addr(":8080"))
}

func weiboApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	weiboApp := mvc.New(app)
	expiresTime, _ := time.ParseDuration("168h")
	sess := sessions.New(sessions.Config{Cookie: "weibo_app_cookie", Expires: expiresTime})
	weiboApp.Register(
		sess.Start,
	)
	helloWorldController := new(controllers.HelloWorldController)
	usersController := new(controllers.UsersController)
	weiboApp.Handle(helloWorldController)
	weiboApp.Handle(usersController)
	return app
}
