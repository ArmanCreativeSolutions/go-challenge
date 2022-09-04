package usersegments

import (
	"github.com/google/uuid"
	"time"
)

type UserSegment struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Segment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
