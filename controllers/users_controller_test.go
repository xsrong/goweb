package controllers

import (
	"goweb/models"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

func setCtxRequest(ctx iris.Context, body io.ReadCloser, contentLength int64) {
	newRequest, _ := http.NewRequest("FAKEMETHOD", "/fake_path", nil)
	newRequest.Body = body
	newRequest.ContentLength = contentLength
	ctx.ResetRequest(newRequest)
}

func setCtxResponse(ctx iris.Context) {
	w := context.AcquireResponseWriter()
	hw := httptest.NewRecorder()
	w.BeginResponse(hw)
	ctx.ResetResponseWriter(w)
}

func setSession(ctx iris.Context, controller *UsersController) {
	cookie := http.Cookie{Name: "sample_cookie_uuid", Value: ""}
	ctx.SetCookie(&cookie)
	sess := sessions.New(sessions.Config{Cookie: "weibo_app_cookie", Expires: 120000000000})
	controller.Session = sess.Start(ctx)
}

func TestUserCreate(t *testing.T) {
	ctx := context.NewContext(iris.New())
	file, _ := os.Open("sample_user.json")
	defer file.Close()
	setCtxRequest(ctx, file, 500)
	controller := UsersController{}
	user, err := controller.Create(ctx)

	if *user.Email != "email1@sample.com" {
		t.Errorf("Email is expected to be \"email1@sample.com\" but got \"%s\"\n", *user.Email)
	}

	if *user.Password != models.Encrypt("password1") {
		t.Errorf("Password is expected to be \"%s\" but got \"%s\"\n", models.Encrypt("password1"), *user.Password)
	}

	if *user.Username != "username1" {
		t.Errorf("Username is expected to be \"username1\" but got \"%s\"\n", *user.Username)
	}

	if user.Message != "message1" {
		t.Errorf("Message is expected to be \"message1\" but got \"%s\"\n", user.Message)
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
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)

	file, _ := os.Open("sample_login_user.json")
	defer file.Close()
	setCtxRequest(ctx, file, 500)

	controller := UsersController{}
	setSession(ctx, &controller)
	user, err := controller.Login(ctx)

	if err != nil {
		t.Error("expected no error, but an error occured:", err)
	}
	if user.ID != 1 {
		t.Errorf("expected returned user id to be 1, but got %d\n:", user.ID)
	}
	id, _ := controller.Session.GetInt("userID")
	if id != 1 {
		t.Errorf("expected user id in session to be 1, but got %d\n", id)
	}
	if crtUsrID := CurrentUser(ctx, controller.Session).ID; crtUsrID != 1 {
		t.Errorf("expected current user id to be 1, but got %d\n", crtUsrID)
	}
}

func TestUserLogout(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	err := controller.Logout()
	if err.Error() != "not logged in" {
		t.Errorf("expected get \"not logged in\" but got \"%s\"\n", err.Error())
	}

	controller.Session.Set("userID", 1)
	err = controller.Logout()
	userID, _ := controller.Session.GetInt("userID")
	if userID >= 0 {
		t.Errorf("expected user id in session to be -1 but got %d\n", userID)
	}
	if err != nil {
		t.Error("got an error:", err)
	}
}

func TestUserUpdate(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)

	file, _ := os.Open("update_user.json")
	defer file.Close()
	setCtxRequest(ctx, file, 500)

	controller := UsersController{}
	setSession(ctx, &controller)
	user, err := controller.Update(1, ctx)

	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	controller.Session.Set("userID", 1)
	user, err = controller.Update(1, ctx)

	if *user.Password != models.Encrypt("newPassword") {
		t.Errorf("Password is expected to be \"%s\" but got \"%s\"\n", models.Encrypt("newPassword"), *user.Password)
	}

	if *user.Username != "new username" {
		t.Errorf("Username is expected to be \"new username\" but got \"%s\"\n", *user.Username)
	}

	if user.Message != "new message" {
		t.Errorf("Message is expected to be \"new message\" but got \"%s\"\n", user.Message)
	}

	if err != nil {
		t.Error(err)
	}
}
