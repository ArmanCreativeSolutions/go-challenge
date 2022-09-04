package repository

import (
	"errors"
	"github.com/google/uuid"
	"go-challenge/application/user-segments/dto"
	"go-challenge/domain/usersegments"
	"go-challenge/infrastructure/dbconfig"
	"time"
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

func (userSegmentRepository *UserSegmentRepository) RemoveSegmentScheduledResponse() error {
	var userSegment usersegments.UserSegment
	now := time.Now()
	err := userSegmentRepository.dbService.GetDB().Debug().
		Model(&userSegment).
		Where("julianday('now') - julianday('created_at') > ?", 14).
		Update("updated_at", now).
		Error
	return err
}
