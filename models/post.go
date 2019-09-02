package models

import (
	"errors"
	"time"
)

type Post struct {
	ID        int
	Content   *string `gorm:"not null"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) Create() (err error) {
	err = DB.Create(p).Error
	if err != nil {
		err = errors.New("error occured when creating post")
	}
	return
}
