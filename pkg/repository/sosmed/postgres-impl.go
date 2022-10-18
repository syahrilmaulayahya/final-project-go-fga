package sosmed

import (
	"context"

	"github.com/syahrilmaulayahya/final-project-go-fga/config/postgres"
	"github.com/syahrilmaulayahya/final-project-go-fga/pkg/domain/sosmed"
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
