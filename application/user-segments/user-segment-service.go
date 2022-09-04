package user_segments

import (
	"go-challenge/application/user-segments/dto"
	"go-challenge/domain/usersegments"
	"go-challenge/infrastructure/repository"
)

type UserSegmentService struct {
	userSegmentRepository *repository.UserSegmentRepository
}

func NewUserSegmentService(userSegmentRepository *repository.UserSegmentRepository) *UserSegmentService {
	return &UserSegmentService{userSegmentRepository: userSegmentRepository}
}

func (userSegmentService *UserSegmentService) CreateUserSegment(createUserSegmentRequest dto.CreateUserSegmentDTO) (usersegments.UserSegment, error) {
	err := userSegmentService.userSegmentRepository.CheckUserId(createUserSegmentRequest.UserId)
	if err != nil {
		return usersegments.UserSegment{Id: createUserSegmentRequest.UserId}, err
	}
	createUserResponse, createError := userSegmentService.userSegmentRepository.Create(createUserSegmentRequest)
	if createError != nil {
		return usersegments.UserSegment{}, createError
	}
	return createUserResponse, nil
}
