package controllers

import (
	"goweb/models"
	"net/http"
	"os"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func TestUserCreate(t *testing.T) {
	models.DB.Delete(&models.User{})
	ctx := context.NewContext(iris.New())
	file, _ := os.Open("sample_user.json")
	defer file.Close()
	newRequest, _ := http.NewRequest("POST", "users/new", nil)
	newRequest.ContentLength = 500
	newRequest.Body = file
	ctx.ResetRequest(newRequest)

	controller := UsersController{}
	user, err := controller.Create(ctx)

	if *user.Email != "email1@sample.com" {
		t.Errorf("Email is expected to be \"email1@sample.com\" but got \"%s\"\n", *user.Email)
	}

	if *user.Password != models.Encrypt("password1") {
		t.Errorf("Email is expected to be \"%s\" but got \"%s\"\n", models.Encrypt("password1"), *user.Password)
	}

	if *user.Username != "username1" {
		t.Errorf("Email is expected to be \"username1\" but got \"%s\"\n", *user.Username)
	}

	if user.Message != "message1" {
		t.Errorf("Email is expected to be \"message1\" but got \"%s\"\n", user.Message)
	}

	if err != nil {
		t.Error(err)
	}

}

func TestFindUserByID(t *testing.T) {
	id := 1
	controller := UsersController{}
	user, err := controller.Show(id)
	if user.ID != 1 || err != nil {
		t.Error("expected to show user but error occured:", user, err)
	}
}

func TestUserLogin(t *testing.T) {
	ctx := context.NewContext(iris.New())
	file, _ := os.Open("sample_login_user.json")
	defer file.Close()
	newRequest, _ := http.NewRequest("POST", "/login", nil)
	newRequest.ContentLength = 500
	newRequest.Body = file
	ctx.ResetRequest(newRequest)

	controller := UsersController{}
	controller.Login(ctx)
	id, _ := controller.Session.GetInt("id")
	if id != 1 {
		t.Errorf("expectd user id in session to be 1, but got %d", id)
	}
}
