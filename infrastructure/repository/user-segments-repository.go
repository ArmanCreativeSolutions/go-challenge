package repository

import (
	"errors"
	"github.com/google/uuid"
	"go-challenge/application/user-segments/dto"
	"go-challenge/domain/usersegments"
	"go-challenge/infrastructure/dbconfig"
)

type UserSegmentRepository struct {
	dbService *dbconfig.SqliteDBService
}

func NewUserSegmentRepository(dbService *dbconfig.SqliteDBService) *UserSegmentRepository {
	return &UserSegmentRepository{dbService}
}

func (userSegmentRepository *UserSegmentRepository) CheckUserId(userId uuid.UUID) error {
	var user usersegments.UserSegment
	var exists int64
	userSegmentRepository.dbService.GetDB().
		Model(&user).
		Count(&exists).
		Where("id = ?", userId)
	if exists > 0 {
		return errors.New("user already exist")
	}
	return nil
}

func (userSegmentRepository *UserSegmentRepository) Create(dto dto.CreateUserSegmentDTO) (usersegments.UserSegment, error) {
	userSegment := usersegments.UserSegment{
		Id:      dto.UserId,
		Segment: dto.Segment,
	}
	result := userSegmentRepository.dbService.GetDB().Create(&userSegment).Error
	return userSegment, result
}

func (userSegmentRepository *UserSegmentRepository) CountUserSegmentsBySegmentTitle(segmentTitle string) (int64, error) {
	var userSegment usersegments.UserSegment
	var userSegmentCount int64
	err := userSegmentRepository.dbService.GetDB().
		Model(&userSegment).
		Count(&userSegmentCount).
		Where("segment = ?", segmentTitle).
		Error
	return userSegmentCount, err
}
