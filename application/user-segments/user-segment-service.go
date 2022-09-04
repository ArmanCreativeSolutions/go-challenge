package user_segments

import "go-challenge/domain/usersegments"

type UserSegmentService struct {
	userSegmentRepository *usersegments.UserSegmentRepository
}

func NewUserSegmentService(userSegmentRepository *usersegments.UserSegmentRepository) *UserSegmentService {
	return &UserSegmentService{userSegmentRepository: userSegmentRepository}
}
