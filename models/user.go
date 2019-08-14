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
	CreatedAt time.Time
	UpdatedAt time.Time
}

var QueryKey string = "id, username, message"

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
