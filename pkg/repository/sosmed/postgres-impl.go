package sosmed

import (
	"context"
	"errors"

	customError "github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/errors"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
	"gorm.io/gorm"
)

type SosmedRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewSosmedRepo(pgCln postgres.PostgresClient) sosmed.SosmedRepo {
	return &SosmedRepoImpl{pgCln: pgCln}
}

func (s *SosmedRepoImpl) AddSosmed(ctx context.Context, userId uint, input *sosmed.Sosmed) error {
	db := s.pgCln.GetClient()
	input.UserID = userId
	result := db.Model(&sosmed.Sosmed{}).Create(&input)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *SosmedRepoImpl) GetSosmedByUserId(ctx context.Context, userId uint) ([]sosmed.Sosmed, error) {
	var result []sosmed.Sosmed
	db := s.pgCln.GetClient()
	resultDb := db.Model(&sosmed.Sosmed{}).Find(&result, "user_id = ?", userId)
	if resultDb.Error != nil {
		return nil, resultDb.Error
	}
	return result, nil
}

func (s *SosmedRepoImpl) UpdateSosmed(ctx context.Context, userId, Id uint, input sosmed.Sosmed) (sosmed.Sosmed, error) {
	var result sosmed.Sosmed
	db := s.pgCln.GetClient()
	resultDb := db.Model(&sosmed.Sosmed{}).First(&result, "user_id = ? AND id = ?", userId, Id)
	if errors.Is(resultDb.Error, gorm.ErrRecordNotFound) {
		return sosmed.Sosmed{}, customError.ErrSosmedNotFound
	}
	if input.Name != "" {
		result.Name = input.Name
	}
	result.Url = input.Url
	if input.Url != "" {
		result.Url = input.Url
	}
	resultDb.Save(&result)
	if resultDb.Error != nil {
		return sosmed.Sosmed{}, resultDb.Error
	}

	return result, nil
}
