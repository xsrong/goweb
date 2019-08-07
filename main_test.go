package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestWeiboApp(t *testing.T) {
	app := weiboApp()
	e := httptest.New(t, app)
	e.GET("/home").Expect().Status(httptest.StatusOK)
}
