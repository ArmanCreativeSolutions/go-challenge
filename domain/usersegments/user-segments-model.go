package usersegments

import (
	"github.com/google/uuid"
	"time"
)

type UserSegment struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Segment   string
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}
