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
