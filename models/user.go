package models

import (
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

func (u *User) Create() (*User, error) {
	err := DB.Create(&u).Error
	return u, err
}
