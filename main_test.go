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
