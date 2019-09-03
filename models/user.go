package models

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        int
	Email     *string `gorm:"not null;unique_index"`
	Password  *string `gorm:"not null"`
	Username  *string `gorm:"not null;unique_index"`
	Message   string
	Following []User `gorm:"many2many:relationships;association_jointable_foreignkey:follow_to"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

var QueryKey string = "users.id, username, message"

func (u *User) Create() (err error) {
	if u.Email != nil && !u.isValidEmail() {
		err = errors.New("Invalid email address")
		return err
	}
	if u.Password == nil {
		err = errors.New("Error occured when creating user")
		return err
	}
	plain := *u.Password
	encrypt := Encrypt(plain)
	u.Password = &encrypt
	err = DB.Create(&u).Error
	if err != nil {
		err = errors.New("Error occured when creating user")
	}
	return
}

func (u *User) isValidEmail() bool {
	pat := `(?i)\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z`
	email := u.Email
	ok, _ := regexp.MatchString(pat, *email)
	return ok
}

func FindUserByID(id int) (user User, err error) {
	DB.Where("id = ?", id).Select(QueryKey).Find(&user)
	if user.ID == 0 {
		err = errors.New("Cannot find such user")
	}
	return
}

func FindUserByEmail(email string) (user User, err error) {
	DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		err = errors.New("Cannot find such user")
	}
	return
}

func FindUserByUsername(username string) (user User, err error) {
	DB.Where("username = ?", username).Select(QueryKey).Find(&user)
	if user.ID == 0 {
		err = errors.New("Cannot find such user")
	}
	return
}

func (u *User) Authenticate() (user User, err error) {
	user, err = FindUserByEmail(*u.Email)
	if err != nil {
		return
	}
	if user.ID == 0 || *user.Password != Encrypt(*u.Password) {
		user = User{}
		err = errors.New("Invalid email or password")
		return
	}
	return
}

func (u *User) Update(params map[string]string) (err error) {
	var password, username *string
	pwd := params["password"]
	if pwd != "" {
		encrypt := Encrypt(pwd)
		password = &encrypt
	}
	usn := params["username"]
	if usn != "" {
		username = &usn
	}
	err = DB.Model(u).Updates(User{Password: password, Username: username, Message: params["message"]}).Error
	if err != nil {
		err = errors.New("Error occured when updating user")
	}
	return
}

// func (u *User) IsFollowed(followTo User) (res bool) {
// 	var rela Relationship
// 	DB.Table("relationships").Where("user_id = ? AND follow_to = ?", u.ID, followTo.ID).Find(&rela)
// 	if rela.ID != 0 {
// 		return true
// 	}
// 	return false
// }

func (u *User) Follow(followTo User) (err error) {
	err = DB.Model(u).Association("Following").Append(&followTo).Error
	if err != nil {
		err = errors.New("Follow failed")
	}
	return
}

func (u *User) Unfollow(followed User) (err error) {
	err = DB.Model(u).Association("Following").Delete(&followed).Error
	if err != nil {
		err = errors.New("Unfollow failed")
	}
	return
}

func (u *User) Followings() (users []User, err error) {
	err = DB.Model(u).Select(QueryKey).Related(&users, "Following").Error
	if err != nil {
		err = errors.New("Error occured when getting following users.")
	}
	return
}

func (u *User) Followers() (users []User, err error) {
	err = DB.Select(QueryKey).Joins("JOIN relationships ON relationships.user_id = users.id").Where("relationships.follow_to = ?", u.ID).Find(&users).Error
	if err != nil {
		err = errors.New("Error occured when getting followers.")
	}
	return
}

func (u *User) IsFollowEachOther(otherUser User) (res bool) {
	var rela Relationship
	DB.Raw("SELECT * FROM relationships WHERE user_id = ? AND follow_to = ? AND EXISTS (SELECT id FROM relationships WHERE (user_id = ? and follow_to = ?))", u.ID, otherUser.ID, otherUser.ID, u.ID).Find(&rela)
	if rela.ID != 0 {
		res = true
	}
	return
}

func (u *User) AddPost(post *Post) (err error) {
	err = post.Create()
	if err != nil {
		return
	}
	err = DB.Model(u).Association("Posts").Append(post).Error
	return
}
