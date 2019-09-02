package models

import (
	"testing"
)

// func setup(filename string) (users []User) {
// 	file, _ := os.Open(filename)
// 	defer file.Close()
// 	data, _ := ioutil.ReadAll(file)
// 	json.Unmarshal(data, &users)
// 	return users
// }

// func TestUserCreate(t *testing.T) {
// 	DB.Delete(&User{})
// 	users := setup("test_user_data.json")
// 	err := users[0].Create() // 测试Email为空的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Error occured when creating user" {
// 		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
// 	}

// 	err = users[1].Create() // 测试Passworld为空的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Error occured when creating user" {
// 		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err)
// 	}

// 	err = users[2].Create() // 测试Username为空的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Error occured when creating user" {
// 		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
// 	}

// 	err = users[3].Create() // 测试Message为空的情况，注：Message允许为空
// 	if err != nil {
// 		t.Errorf("expected no error but an error occured:\n %s\n", err)
// 	}

// 	err = users[4].Create() // 测试都不为空的情况
// 	if err != nil {
// 		t.Errorf("expected no error but an error occured:\n %s\n", err)
// 	}

// 	err = users[5].Create() // 测试Email发生重复的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Error occured when creating user" {
// 		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
// 	}

// 	err = users[6].Create() // 测试Username发生重复的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Error occured when creating user" {
// 		t.Errorf("expected get \"Error occured when creating user\" but got \"%s\"\n", err.Error())
// 	}

// 	err = users[7].Create() // 测试Email为无效的邮箱地址的情况
// 	if err == nil {
// 		t.Error("expected get an error but no error occured.")
// 	} else if err.Error() != "Invalid email address" {
// 		t.Errorf("expected get \"Invalid email address\" but got \"%s\"\n", err.Error())
// 	}
// }

// func TestFindUserByID(t *testing.T) {
// 	id1, id2 := 5, 4
// 	user1, err := FindUserByID(id1)
// 	if err == nil {
// 		t.Error("expected find no user but got:", user1)
// 	}

// 	user2, err := FindUserByID(id2)
// 	if err != nil {
// 		t.Error("got an unexpected error:", err)
// 	} else if user2.ID != id2 {
// 		t.Errorf("expected find user which id is %d but got user which id is: %d\n", id2, user2.ID)
// 	} else if user2.Email != nil || user2.Password != nil {
// 		t.Error("expected do not get user email and password but got them now")
// 	}
// }

// func TestFindUserByEmail(t *testing.T) {
// 	email1, email2 := "email6@exam.com", "email5@exam.com"
// 	user1, err := FindUserByEmail(email1)
// 	if err == nil {
// 		t.Error("expected find no user but got:", user1)
// 	}

// 	user2, err := FindUserByEmail(email2)
// 	if err != nil {
// 		t.Error("got an unexpected error:", err)
// 	} else if *user2.Email != email2 {
// 		t.Errorf("expected find user which email is \"%s\" but got user which email is: \"%s\"", email2, *user2.Email)
// 	}
// }

// func TestFindUserByUsername(t *testing.T) {
// 	username1, username2 := "username6", "username5"
// 	user1, err := FindUserByUsername(username1)
// 	if err == nil {
// 		t.Error("expected find no user but got:", user1)
// 	}

// 	user2, err := FindUserByUsername(username2)
// 	if err != nil {
// 		t.Error("got an unexpected error:", err)
// 	} else if *user2.Username != username2 {
// 		t.Errorf("expected find user which username is \"%s\" but got user which username is: \"%s\"", username2, *user2.Username)
// 	} else if user2.Email != nil || user2.Password != nil {
// 		t.Error("expected do not get user email and password but got them now")
// 	}
// }

// func TestUserAuthenticate(t *testing.T) {
// 	users := setup("test_auth_data.json")

// 	if _, err := users[0].Authenticate(); err == nil {
// 		if err.Error() != "Invalid email or password" {
// 			t.Errorf("expected get \"Invalid email or password\" but got \"%s\"\n", err.Error())
// 		}
// 		t.Error("expected authentication fail but it passed")
// 	}

// 	if _, err := users[1].Authenticate(); err == nil {
// 		if err.Error() != "Invalid email or password" {
// 			t.Errorf("expected get \"Invalid email or password\" but got \"%s\"\n", err.Error())
// 		}
// 		t.Error("expected authentication fail but it passed")
// 	}

// 	if user, err := users[2].Authenticate(); err != nil {
// 		t.Error("expected authentication pass but it failed cause:", err)
// 	} else if user.ID != 3 {
// 		t.Errorf("expected user id to be 3 but got %d\n", user.ID)
// 	}
// }

func TestUpdateUser(t *testing.T) {
	oldEmail := "old_email@example.com"
	newEmail := "new_email@example.com"
	oldPassword := "oldpassword"
	newPassword := "newpassword"
	oldUsername := "old username"
	newUsername := "new username"
	oldMessage := "old message"
	newMessage := "new message"

	user := User{Email: &oldEmail, Password: &oldPassword, Username: &oldUsername, Message: oldMessage}
	user.Create()

	// test with params1
	params1 := map[string]string{"password": "", "username": "", "message": ""}
	if err := user.Update(params1); err != nil {
		t.Error("got an error:", err)
	}
	u, _ := FindUserByEmail(oldEmail)
	if *u.Password != Encrypt(oldPassword) || *u.Username != oldUsername || u.Message != oldMessage {
		t.Error("test failed at params1")
	}

	// test with params2
	params2 := map[string]string{"password": newPassword, "username": "", "message": ""}
	if err := user.Update(params2); err != nil {
		t.Error("got an error:", err)
	}
	u, _ = FindUserByEmail(oldEmail)
	if *u.Password != Encrypt(newPassword) || *u.Username != oldUsername || u.Message != oldMessage {
		t.Error("test failed at params2")
	}

	// test with params3
	params3 := map[string]string{"password": newPassword, "username": newUsername, "message": newMessage}
	if err := user.Update(params3); err != nil {
		t.Error("got an error:", err)
	}
	u, _ = FindUserByEmail(oldEmail)
	if *u.Password != Encrypt(newPassword) || *u.Username != newUsername || u.Message != newMessage {
		t.Error("test failed at params3")
	}

	// test with params4
	params4 := map[string]string{"email": newEmail}
	if err := user.Update(params4); err != nil {
		t.Error("got an error:", err)
	}
	u, _ = FindUserByEmail(newEmail)
	if u.ID != 0 {
		t.Error("test failed at params4")
	}

}

func TestFollowUser(t *testing.T) {
	e1 := "111@11.com"
	p1 := "11111111"
	u1 := "11111"
	e2 := "222@22.com"
	p2 := "22222222"
	u2 := "22222"
	user1 := User{Email: &e1, Password: &p1, Username: &u1}
	user2 := User{Email: &e2, Password: &p2, Username: &u2}
	user1.Create()
	user2.Create()
	user1.Follow(user2)
	var rela Relationship
	DB.Table("relationships").Where("user_id = ? AND follow_to = ?", user1.ID, user2.ID).Find(&rela)
	if rela.ID == 0 {
		t.Error("test follow user failed")
	}
}
