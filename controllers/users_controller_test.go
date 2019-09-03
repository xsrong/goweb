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
	// if crtUsrID := CurrentUser(ctx, controller.Session).ID; crtUsrID != 1 {
	// 	t.Errorf("expected current user id to be 1, but got %d\n", crtUsrID)
	// }

	_, err = controller.Login(ctx)

	if err.Error() != "logged in already" {
		t.Errorf("expected get \"logged in already\" but got \"%s\"\n", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	// 测试用户未登录时调用UersController#Logout()
	err := controller.Logout()
	if err.Error() != "not logged in" {
		t.Errorf("expected get \"not logged in\" but got \"%s\"\n", err.Error())
	}

	// 测试用户登录后调用UersController#Logout()
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

	// 测试未登录时更新用户信息
	controller := UsersController{}
	setSession(ctx, &controller)
	user, err := controller.Update(1, ctx)

	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试用户已登录，但更新的用户信息非已登录的用户时的情况
	controller.Session.Set("userID", 1)
	user, err = controller.Update(2, ctx)

	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试用户已登录且更新的用户信息为该用户时的情况
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

func TestUserFlollow(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	email := "email2@sample.com"
	password := "password2"
	username := "username2"
	anotherUser := models.User{Email: &email, Password: &password, Username: &username}
	anotherUser.Create()

	// 测试用户未登录时关注其他用户
	err := controller.Follow(2, 1)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试发起关注的用户不是登录用户的情况
	controller.Session.Set("userID", 1)
	err = controller.Follow(2, 1)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试用户关注自己的情况
	controller.Session.Set("userID", 2)
	err = controller.Follow(2, 2)
	if err.Error() != "cannot follow yourself" {
		t.Errorf("expected \"cannot follow yourself\" but got \"%s\"\n", err.Error())
	}

	// 测试正常关注的情况
	controller.Session.Set("userID", 2)
	err = controller.Follow(2, 1)
	if err != nil {
		t.Errorf("expected no error but got \"%s\"\n", err.Error())
	}
}

func TestUserUnfollow(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	// 测试用户未登录时取消关注其他用户
	err := controller.Unfollow(2, 1)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试发起取消关注的用户不是登录用户的情况
	controller.Session.Set("userID", 1)
	err = controller.Unfollow(2, 1)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试用户对自己取消关注的情况
	controller.Session.Set("userID", 2)
	err = controller.Unfollow(2, 2)
	if err.Error() != "cannot unfollow yourself" {
		t.Errorf("expected \"cannot unfollow yourself\" but got \"%s\"\n", err.Error())
	}

	// 测试正常取消关注的情况
	controller.Session.Set("userID", 2)
	err = controller.Unfollow(2, 1)
	if err != nil {
		t.Errorf("expected no error but got \"%s\"\n", err.Error())
	}
}

func setupUserFollowingAndFollowers() {
	user1, _ := models.FindUserByID(1)
	user2, _ := models.FindUserByID(2)
	user3, _ := models.FindUserByID(3)

	user1.Follow(user2)
	user1.Follow(user3)
	user2.Follow(user3)
}

func TestGetUserFollowing(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	email := "email3@sample.com"
	password := "password3"
	username := "username3"
	anotherUser := models.User{Email: &email, Password: &password, Username: &username}
	anotherUser.Create()

	setupUserFollowingAndFollowers()

	// 测试未登录时请求全部关注对象
	_, err := controller.Following(1)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试请求当前登录用户以外的其他用户的关注对象
	controller.Session.Set("userID", 1)
	_, err = controller.Following(2)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试正常请求用户的关注对象
	_, err = controller.Following(1)
	if err != nil {
		t.Errorf("expected no error but got \"%s\"\n", err.Error())
	}
}

func TestGetUserFollowers(t *testing.T) {
	app := iris.New()
	ctx := context.NewContext(app)
	setCtxResponse(ctx)
	setCtxRequest(ctx, nil, 0)
	controller := UsersController{}
	setSession(ctx, &controller)

	// 测试未登录时请求被谁关注
	_, err := controller.Followers(3)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试请求当前登录用户以外的其他用户的被谁关注
	controller.Session.Set("userID", 3)
	_, err = controller.Followers(2)
	if err.Error() != "authenticate failed! please login and try again" {
		t.Errorf("expected \"authenticate failed! please login and try again\" but got \"%s\"\n", err.Error())
	}

	// 测试正常请求当前用户被谁关注
	_, err = controller.Followers(3)
	if err != nil {
		t.Errorf("expected no error but got \"%s\"\n", err.Error())
	}
}
