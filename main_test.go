package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestWeiboApp(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)
	e.GET("/home").Expect().Status(httptest.StatusOK).Body().Equal("Hello World!")
}

func TestUsersCreateRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/users/new")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "password1", "username": "username1", "message": "message1"})
	request.Expect().Status(httptest.StatusOK)
}

func TestUsersShowRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("GET", "/users/1")
	response := request.Expect()
	response.Status(httptest.StatusOK)
}

func TestLoginRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "password1"})

	response := request.Expect()
	response.Status(httptest.StatusOK)
}

func TestLogoutRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "password1"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("DELETE", "/logout")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect().Status(httptest.StatusOK)
}

func TestUserUpdateRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "password1"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("PATCH", "/users/1/edit")
	request.WithJSON(map[string]interface{}{"password": "newPassword", "username": "new username", "message": "new message"})
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	response := request.Expect()
	response.Status(httptest.StatusOK)
}

func TestFollowUserRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/users/new")
	request.WithJSON(map[string]interface{}{"email": "email2@example.com", "password": "password2", "username": "username2", "message": "message2"})
	request.Expect()

	request = e.Request("POST", "/users/new")
	request.WithJSON(map[string]interface{}{"email": "email3@example.com", "password": "password3", "username": "username3", "message": "message3"})
	request.Expect()

	request = e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "newPassword"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("POST", "/users/1/follow/2")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect().Status(httptest.StatusOK)
}

func TestUnfollowUserRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "newPassword"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("DELETE", "/users/1/unfollow/2")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect().Status(httptest.StatusOK)
}

func TestGetFollowingRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email1@example.com", "password": "newPassword"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("POST", "/users/1/follow/2")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect()

	request = e.Request("POST", "/users/1/follow/3")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect()

	request = e.Request("GET", "/users/1/following")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect().Status(httptest.StatusOK)
}

func TestGetFollowersRoute(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)

	request := e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email2@example.com", "password": "password2"})
	cookie := request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("POST", "/users/2/follow/3")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect()

	request = e.Request("POST", "/login")
	request.WithJSON(map[string]interface{}{"email": "email3@example.com", "password": "password3"})
	cookie = request.Expect().Cookie("weibo_app_cookie")

	request = e.Request("GET", "/users/3/followers")
	request.WithCookie(cookie.Name().Raw(), cookie.Value().Raw())
	request.Expect().Status(httptest.StatusOK)

}
