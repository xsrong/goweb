package models

import (
	"errors"
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

func (u *User) Create() error {
	plain := *u.Password
	encrypt := Encrypt(plain)
	u.Password = &encrypt
	err := DB.Create(&u).Error
	if err != nil {
		err = errors.New("Error occured when creating user.")
	}
	return err
}
