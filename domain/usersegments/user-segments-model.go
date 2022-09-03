package usersegments

import (
	"gorm.io/gorm"
	"time"
)

type UserSegment struct {
	gorm.Model
	Id        string
	Segment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
