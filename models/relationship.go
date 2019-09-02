package models

import "time"

type Relationship struct {
	ID        int
	UserID    int
	FollowTo  int
	CreatedAt time.Time
}
