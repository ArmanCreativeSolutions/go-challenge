package dto

type CreateUserSegmentDTO struct {
	UserId  string `validate:"required"`
	Segment string `validate:"required"`
}
