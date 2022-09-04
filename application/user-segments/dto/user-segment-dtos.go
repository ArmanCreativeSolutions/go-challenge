package dto

import "github.com/google/uuid"

type CreateUserSegmentDTO struct {
	UserId  uuid.UUID `json:"userId" validate:"required"`
	Segment string    `json:"segment" validate:"required"`
}
