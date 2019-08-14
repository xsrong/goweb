package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func setup(filename string) (users []User) {
	file, _ := os.Open(filename)
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &users)
	return users
}

func TestUserCreate(t *testing.T) {
	DB.Delete(&User{})
	users := setup("test_user_data.json")
	err := users[0].Create() // 测试Email为空的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Error occured when creating user" {
		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
	}

	err = users[1].Create() // 测试Passworld为空的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Error occured when creating user" {
		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err)
	}

	err = users[2].Create() // 测试Username为空的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Error occured when creating user" {
		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
	}

	err = users[3].Create() // 测试Message为空的情况，注：Message允许为空
	if err != nil {
		t.Errorf("expected no error but an error occured:\n %s\n", err)
	}

	err = users[4].Create() // 测试都不为空的情况
	if err != nil {
		t.Errorf("expected no error but an error occured:\n %s\n", err)
	}

	err = users[5].Create() // 测试Email发生重复的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Error occured when creating user" {
		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
	}

	err = users[6].Create() // 测试Username发生重复的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Error occured when creating user" {
		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
	}

	err = users[7].Create() // 测试Email为无效的邮箱地址的情况
	if err == nil {
		t.Error("expected get an error but no error occured.")
	} else if err.Error() != "Invalid email address" {
		t.Errorf("expected get \"Invalid email address\" but got \"%s\"\n", err.Error())
	}
}

func TestFindUserByID(t *testing.T) {
	id1, id2 := 5, 4
	user1, err := FindUserByID(id1)
	if err == nil {
		t.Error("expected find no user but got:", user1)
	}

	user2, err := FindUserByID(id2)
	if err != nil {
		t.Error("got an unexpected error:", err)
	} else if user2.ID != id2 {
		t.Errorf("expected find user which id is %d but got user which id is: %d\n", id2, user2.ID)
	} else if user2.Email != nil || user2.Password != nil {
		t.Error("expected do not get user email and password but got them now")
	}
}

func TestFindUserByEmail(t *testing.T) {
	email1, email2 := "email6@exam.com", "email5@exam.com"
	user1, err := FindUserByEmail(email1)
	if err == nil {
		t.Error("expected find no user but got:", user1)
	}

	user2, err := FindUserByEmail(email2)
	if err != nil {
		t.Error("got an unexpected error:", err)
	} else if *user2.Email != email2 {
		t.Errorf("expected find user which email is \"%s\" but got user which email is: \"%s\"", email2, *user2.Email)
	}
}

func TestFindUserByUsername(t *testing.T) {
	username1, username2 := "username6", "username5"
	user1, err := FindUserByUsername(username1)
	if err == nil {
		t.Error("expected find no user but got:", user1)
	}

	user2, err := FindUserByUsername(username2)
	if err != nil {
		t.Error("got an unexpected error:", err)
	} else if *user2.Username != username2 {
		t.Errorf("expected find user which username is \"%s\" but got user which username is: \"%s\"", username2, *user2.Username)
	} else if user2.Email != nil || user2.Password != nil {
		t.Error("expected do not get user email and password but got them now")
	}
}

func TestUserAuthenticate(t *testing.T) {
	users := setup("test_auth_data.json")

	if _, err := users[0].Authenticate(); err == nil {
		if err.Error() != "Invalid email or password" {
			t.Errorf("expected get \"Invalid email or password\" but got \"%s\"\n", err.Error())
		}
		t.Error("expected authentication fail but it passed")
	}

	if _, err := users[1].Authenticate(); err == nil {
		if err.Error() != "Invalid email or password" {
			t.Errorf("expected get \"Invalid email or password\" but got \"%s\"\n", err.Error())
		}
		t.Error("expected authentication fail but it passed")
	}

	if user, err := users[2].Authenticate(); err != nil {
		t.Error("expected authentication pass but it failed cause:", err)
	} else if user.ID != 3 {
		t.Errorf("expected user id to be 3 but got %d\n", user.ID)
	}
}
