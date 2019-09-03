package models

import "time"

type Relationship struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	FollowTo  int
	CreatedAt time.Time
}
